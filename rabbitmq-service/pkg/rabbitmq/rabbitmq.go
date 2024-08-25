package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectToRabbitMQ() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, fmt.Errorf("failed to connect Rabbit MQ: %w", err)
	}

	return conn, nil
}
