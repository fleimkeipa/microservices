package main

import (
	"log"
	"rabbitmq-service/pkg/rabbitmq"
)

func main() {
	// Get Rabbit MQ connection
	conn, err := rabbitmq.ConnectToRabbitMQ()
	if err != nil {
		log.Fatalf("failed to connect Rabbit MQ: %v", err)
	}
	defer conn.Close()

	// create new channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to connect channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"order.created", // name
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		log.Fatalf("failed to control queue: %v", err)
	}

	// Start consumer and consume "order.created"
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	var forever chan struct{}
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
