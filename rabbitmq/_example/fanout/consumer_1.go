package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"

	rabbitmq "github.com/iamgoangle/pkg/rabbitmq"
	middlewares "github.com/iamgoangle/pkg/rabbitmq/middlewares"
)

const (
	exchangeName = "exc_fanout"

	queueName = "q_fanout_1"
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

	connection.Use(middlewares.ExchangeDeclare(exchangeName, middlewares.ExchangeFanout, nil))
	connection.Use(middlewares.QueueDeclare(queueName, nil))
	connection.Use(middlewares.QueueBind(queueName, "", exchangeName, false, nil))
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
