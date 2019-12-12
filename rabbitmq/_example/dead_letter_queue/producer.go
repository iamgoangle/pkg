package main

import (
	"log"

	rabbitmq "github.com/iamgoangle/pkg/rabbitmq"
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

	producer := rabbitmq.NewProducer("exc_main", "", "", connection)
	err = producer.Publish([]byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`))
	if err != nil {
		log.Println("unable to publish body")
	}
}
