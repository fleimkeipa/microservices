package main

import (
	"nats-listener-service/internal/listener"
	"nats-listener-service/pkg/nats"
)

func main() {
	// NATS bağlantısını başlat
	nc := nats.ConnectToNATS()
	defer nc.Close()

	// ListenerService'i başlat ve "order.created" konusunu dinle
	subject := "order.created"
	listenerService := listener.NewListenerService(nc, subject)
	listenerService.Listen()
}
