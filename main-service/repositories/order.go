package repositories

import (
	"log"

	"order-service/repositories/interfaces"
)

type OrderRepository struct {
	natsRepo interfaces.MessageInterfaces
}

func NewOrderRepository(natsRepo interfaces.MessageInterfaces) *OrderRepository {
	return &OrderRepository{
		natsRepo: natsRepo,
	}
}

func (o *OrderRepository) Create(orderID string) error {
	// Create order
	if err := o.natsRepo.Send("order.created", orderID); err != nil {
		log.Printf("Failed to publish order.created: %v", err)
		return err
	}

	return nil
}
