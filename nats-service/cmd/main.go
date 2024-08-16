package main

import (
	"nats-listener-service/internal/listener"
	"nats-listener-service/pkg/nats"
)

func main() {
	// Get NATS connection
	nc := nats.ConnectToNATS()
	defer nc.Close()

	// Start ListenerService and listen "order.created"
	subject := "order.created"
	listenerService := listener.NewListenerService(nc, subject)
	listenerService.Listen()
}
