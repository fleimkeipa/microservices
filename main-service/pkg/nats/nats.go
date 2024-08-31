package nats

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

func ConnectToNATS() (*nats.Conn, error) {
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		return nil, fmt.Errorf("failed to connect NATS: %w", err)
	}

	return nc, nil
}
