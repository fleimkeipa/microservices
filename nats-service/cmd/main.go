package main

import (
	"log"

	"nats-listener-service/internal/listener"
	"nats-listener-service/pkg/nats"
)

func main() {
	// Get NATS connection
	nc, err := nats.ConnectToNATS()
	if err != nil {
		log.Fatalf("failed to connect NATS: %v", err)
		return
	}
	defer nc.Close()

	// Start ListenerService and listen "order.created"
	subject := "order.created"
	listenerService := listener.NewListenerService(nc, subject)
	listenerService.Listen()
}
