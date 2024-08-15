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
	// Belirli bir NATS konusunu (subject) dinle
	_, err := s.nc.Subscribe(s.subject, func(msg *nats.Msg) {
		log.Printf("Received message on subject %s: %s", s.subject, string(msg.Data))

		// Burada gelen mesajı işleyin
		// Örneğin, bir işlemi tetikleyebilir, veritabanına kaydedebilir, vb.
		s.handleMessage(msg.Data)
	})
	if err != nil {
		log.Fatalf("Error subscribing to subject %s: %v", s.subject, err)
	}

	// NATS dinleyicisi açık kalmalı
	select {}
}

func (s *ListenerService) handleMessage(data []byte) {
	// Mesaj verisini işleme
	log.Printf("Processing message: %s", string(data))
	// İşlem gerçekleştir (örneğin, bir veritabanı güncellemesi)
}
