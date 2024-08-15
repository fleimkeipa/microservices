package order

import (
	"context"
	"log"

	"github.com/nats-io/nats.go"
)

type OrderService struct {
	nc *nats.Conn
}

func NewOrderService(nc *nats.Conn) *OrderService {
	return &OrderService{nc: nc}
}

func (s *OrderService) CreateOrder(ctx context.Context, orderID string) error {
	// Sipariş oluşturma işlemi (veritabanına kaydetme vb.)
	log.Printf("Creating order with ID: %s", orderID)

	// İşlem tamamlandığında NATS üzerinden bir mesaj gönder
	err := s.nc.Publish("order.created", []byte(orderID))
	if err != nil {
		log.Printf("Failed to publish order.created: %v", err)
		return err
	}
	return nil
}
