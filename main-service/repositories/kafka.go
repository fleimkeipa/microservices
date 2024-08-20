package repositories

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

type KafkaRepository struct {
	nc *nats.Conn
}

func NewKafkaRepository(nc *nats.Conn) *KafkaRepository {
	return &KafkaRepository{
		nc: nc,
	}
}

func (rc *KafkaRepository) Send(subj string, data string) error {
	// Create order
	if err := rc.nc.Publish(subj, []byte(data)); err != nil {
		return fmt.Errorf("failed to publish, subject: [%s], data: [%s], error: %w", subj, data, err)
	}

	return nil
}
