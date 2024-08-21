package main

import (
	"log"

	"kafka-consumer-service/internal/consservice"
	"kafka-consumer-service/pkg/kafka"
)

func main() {
	// Get Kafka consumer

	// Create a new consumer and start it.
	consumer, err := kafka.ConnectConsumer()
	if err != nil {
		log.Fatalf("failed to connect Kafka: %v", err)
		return
	}
	defer consumer.Close()

	var topic = "order.created"

	var consService = consservice.NewConsumerService(consumer, topic)
	consService.Listen(topic)
}
