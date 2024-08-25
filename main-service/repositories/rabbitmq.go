package repositories

import (
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQRepository struct {
	conn *amqp.Connection
}

func NewRabbitMQRepository(conn *amqp.Connection) *RabbitMQRepository {
	return &RabbitMQRepository{
		conn: conn,
	}
}

func (rc *RabbitMQRepository) Send(subj string, data string) error {
	ch, err := rc.conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to connect channel: %w", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		subj,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to queue declaretion: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(data),
		})

	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}
