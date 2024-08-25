package main

import (
	"log"

	"rabbitmq-service/internal/consumer"
	"rabbitmq-service/pkg/rabbitmq"
)

func main() {
	// Get Rabbit MQ connection
	conn, err := rabbitmq.ConnectToRabbitMQ()
	if err != nil {
		log.Fatalf("failed to connect Rabbit MQ: %v", err)
	}
	defer conn.Close()

	subject := "order.created"
	consumerService := consumer.NewConsumerService(conn, subject)
	consumerService.Consume()
}
