package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"

	rabbitmq "github.com/iamgoangle/pkg/rabbitmq"
	middlewares "github.com/iamgoangle/pkg/rabbitmq/middlewares"
)

const (
	exchangeName = "exc_topic"
	queueName    = "q_topic_broadcast"
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

	connection.Use(middlewares.ExchangeDeclare(exchangeName, middlewares.ExchangeTopic, nil))
	connection.Use(middlewares.QueueDeclare(queueName, nil))
	connection.Use(middlewares.QueueBind(queueName, "campaign.broadcast.*", exchangeName, false, nil))
	if err := connection.Run(); err != nil {
		log.Panic(err)
	}

	consumer := rabbitmq.NewConsumer(queueName, "consumer_broadcast", connection)
	consumer.Use(newConsumerHandler())

	if err := consumer.Consume(); err != nil {
		log.Panic(err)
	}
}

func newConsumerHandler() rabbitmq.ConsumerHandler {
	return &handler{}
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
