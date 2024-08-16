package listener

import (
	"log"

	"github.com/nats-io/nats.go"
)

type ListenerService struct {
	nc      *nats.Conn
	subject string
}

func NewListenerService(nc *nats.Conn, subject string) *ListenerService {
	return &ListenerService{
		nc:      nc,
		subject: subject,
	}
}

func (s *ListenerService) Listen() {
	// Listen specify NATS subject
	_, err := s.nc.Subscribe(s.subject, func(msg *nats.Msg) {
		log.Printf("Received message on subject %s: %s", s.subject, string(msg.Data))

		// Process the message
		s.handleMessage(msg.Data)
	})
	if err != nil {
		log.Fatalf("Error subscribing to subject %s: %v", s.subject, err)
	}

	// Must be open
	select {}
}

func (s *ListenerService) handleMessage(data []byte) {
	// Process the message
	log.Printf("Processing message: %s", string(data))
}
