package retry

import (
	rabbitmq "github.com/iamgoangle/pkg/rabbitmq"
)

// Config retry queue middleware config
type Config struct {
	QueueName string
	TTL       int
}

// RetryQueue ability to DLQ when NACK the message
func RetryQueue(c *Config) rabbitmq.HandlerFunc {
	return func(c rabbitmq.Connection) error {

		return nil
	}

	return nil
}
