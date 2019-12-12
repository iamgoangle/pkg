package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"

	rabbitmq "github.com/iamgoangle/pkg/rabbitmq"
	middlewares "github.com/iamgoangle/pkg/rabbitmq/middlewares"
)

const (
	exchangeName      = "exc_main"
	exchangeDelayName = "exc_delay"

	queueName      = "q_main"
	queueDelayName = "q_delay"
	delayTime      = 60000
)

type handler struct {
}

func main() {
	conf := rabbitmq.Config{
		Host:     "localhost",
		Port:     5672,
		Username: "admin",
		Password: "1234",
		Vhost:    "/",
	}
	connection, err := rabbitmq.NewAMQPConnection(conf)
	if err != nil {
		log.Panic(err)
	}

	// main exchange and queue
	connection.Use(middlewares.ExchangeDeclare(exchangeName, middlewares.ExchangeDirect, nil))
	connection.Use(middlewares.QueueDeclare(queueName, nil))
	connection.Use(middlewares.QueueBind(queueName, "", exchangeName, false, nil))
	if err := connection.Run(); err != nil {
		log.Panic(err)
	}

	// delay exchange and queue
	connection.Use(middlewares.ExchangeDeclare(exchangeDelayName, middlewares.ExchangeDirect, nil))
	argsDelayQueue := make(amqp.Table)
	argsDelayQueue["x-dead-letter-exchange"] = exchangeName
	// argsDelayQueue["x-dead-letter-routing-key"] = DLXRoutingKey
	argsDelayQueue["x-message-ttl"] = delayTime
	connection.Use(middlewares.QueueDeclare(queueDelayName, argsDelayQueue))
	connection.Use(middlewares.QueueBind(queueDelayName, "", exchangeDelayName, false, nil))
	if err := connection.Run(); err != nil {
		log.Panic(err)
	}

	consumer := rabbitmq.NewConsumer(queueName, "", connection)
	consumer.Use(newConsumerHandler())

	if err := consumer.Consume(); err != nil {
		log.Panic(err)
	}
}

func newConsumerHandler() rabbitmq.ConsumerHandler {
	return &handler{}
}

func (h *handler) OnDelay(m amqp.Delivery) error {
	return nil
}

func (h *handler) Do(msg []byte) error {
	fmt.Println(string(msg))

	// return errors.New("error occurs")
	return nil
}

func (h *handler) OnSuccess(m amqp.Delivery) error {
	log.Println("consume item success")

	return nil
}

func (h *handler) OnError(m amqp.Delivery, err error) {
	log.Panic("business process error")
}
