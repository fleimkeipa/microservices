package consumer

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type ConsumerService struct {
	conn    *amqp.Connection
	subject string
}

func NewConsumerService(conn *amqp.Connection, subject string) *ConsumerService {
	return &ConsumerService{
		conn:    conn,
		subject: subject,
	}
}

func (rc *ConsumerService) Consume() {
	// create new channel
	ch, err := rc.conn.Channel()
	if err != nil {
		log.Fatalf("failed to connect channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		rc.subject, // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
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
