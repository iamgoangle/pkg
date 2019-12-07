package main

import (
	"log"

	rabbitmq "github.com/iamgoangle/pkg/rabbitmq"
)

const (
	exchangeName = "exc_fanout"

	queueName = "q_fanout"
)

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

	// connection.Use(middlewares.ExchangeDeclare(exchangeName, middlewares.ExchangeFanout, nil))
	// connection.Use(middlewares.QueueDeclare(queueName, nil))
	// connection.Use(middlewares.QueueBind(queueName, "", exchangeName, false, nil))
	// if err := connection.Run(); err != nil {
	// 	log.Panic(err)
	// }

	producer := rabbitmq.NewProducer(exchangeName, "", "", connection)
	err = producer.Publish([]byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`))
	if err != nil {
		log.Println("unable to publish body")
	}
}
