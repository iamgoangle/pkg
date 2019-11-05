package rabbitmq

//go:generate mockgen -source=./producer.go -destination=./mocks/producer_mock.go -package=mocks github.com/iamgoangle/pkg/rabbitmq Producer

import (
	"github.com/pkg/errors"

	"github.com/streadway/amqp"
)

// Producer handles interface to publish the message to broker and queue
type Producer interface {
	// UseWithConfig configs producer
	UseWithConfig(config ...ProducerConfigHandler) error

	// Publish send the message to amqp broker
	// default is []byte with mime type text/json
	Publish(body []byte) error
}

// Produce represents produce object
type Produce struct {
	exchange, key, kind  string
	mandatory, immediate bool
	conn                 Connection
}

// ProducerConfigHandler handles options parameter as a function
// See https://godoc.org/github.com/streadway/amqp#Channel.Publish
type ProducerConfigHandler func(*Produce) error

// PublishConfigHandler handles pass the publish config func
// See https://godoc.org/github.com/streadway/amqp#Publishing
type PublishConfigHandler func(*amqp.Publishing) error

// NewProducer instance the new producer
func NewProducer(exName, routingKey, kind string, c Connection) Producer {
	return &Produce{
		exchange: exName,
		key:      routingKey,
		kind:     kind,
		conn:     c,
	}
}

func (p *Produce) UseWithConfig(configs ...ProducerConfigHandler) error {
	for _, config := range configs {
		err := config(p)
		if err != nil {
			return errors.Wrap(err, FailedToApplyProducerConfigFunc)
		}
	}

	return nil
}

func (p *Produce) Publish(body []byte) error {
	msg := amqp.Publishing{
		ContentType: "text/json",
		Body:        body,
	}

	// for _, config := range configs {
	// 	err := config(&msg)
	// 	if err != nil {
	// 		return errors.Wrap(err, FailedToApplyPublishPropertyConfigFunc)
	// 	}
	// }

	return p.conn.Publish(p.exchange, p.key, p.mandatory, p.immediate, msg)
}

// PublisherTTLConfig defines closure function to handler publish ttl
func PublisherTTLConfig(ttl string) PublishConfigHandler {
	return func(msg *amqp.Publishing) error {
		if len(ttl) == 0 {
			return errors.New(MissingArgumentTTL)
		}

		msg.Expiration = ttl

		return nil
	}
}

// PublisherDeliveryModeConfig defines closure function to handler delivery mode
// Transient (0 or 1) or Persistent (2)
func PublisherDeliveryModeConfig(persist uint8) PublishConfigHandler {
	return func(msg *amqp.Publishing) error {
		if persist > 2 {
			return errors.New(FailedToSetConfigPersist)
		}

		msg.DeliveryMode = persist

		return nil
	}
}

// PublisherContentTypeConfig set MIME content type
func PublisherContentTypeConfig(cType string) PublishConfigHandler {
	return func(msg *amqp.Publishing) error {
		if len(cType) == 0 {
			return errors.New(MissingArgumentContentType)
		}

		msg.ContentType = cType

		return nil
	}
}

// PublisherPriorityConfig set priority queue
// level 0-9
func PublisherPriorityConfig(level uint8) PublishConfigHandler {
	return func(msg *amqp.Publishing) error {
		if level > 9 {
			return errors.New(FailedToSetConfigPriority)
		}

		msg.Priority = level

		return nil
	}
}
