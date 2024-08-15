package nats

import (
	"log"

	"github.com/nats-io/nats.go"
)

func ConnectToNATS() *nats.Conn {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Error connecting to NATS: %v", err)
	}
	return nc
}
