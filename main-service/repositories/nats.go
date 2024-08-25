package repositories

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

type NATSRepository struct {
	nc *nats.Conn
}

func NewNATSRepository(nc *nats.Conn) *NATSRepository {
	return &NATSRepository{
		nc: nc,
	}
}

func (rc *NATSRepository) Send(subj string, data string) error {
	// Publish message to subject
	if err := rc.nc.Publish(subj, []byte(data)); err != nil {
		return fmt.Errorf("failed to publish, subject: [%s], data: [%s], error: %w", subj, data, err)
	}

	return nil
}
