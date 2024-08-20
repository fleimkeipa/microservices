package commands

import (
	"order-service/repositories/interfaces"
)

type OrderCommandHandlers struct {
	repo interfaces.OrderInterfaces
}

func NewOrderCommandHandlers(repo interfaces.OrderInterfaces) *OrderCommandHandlers {
	return &OrderCommandHandlers{
		repo: repo,
	}
}

func (o *OrderCommandHandlers) Create(orderID string) error {
	return o.repo.Create(orderID)
}
