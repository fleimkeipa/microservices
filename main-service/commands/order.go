package commands

import (
	"order-service/repositories/interfaces"

	"github.com/labstack/echo/v4"
)

type OrderCommandHandlers struct {
	repo interfaces.OrderInterfaces
}

func NewOrderCommandHandlers(repo interfaces.OrderInterfaces) *OrderCommandHandlers {
	return &OrderCommandHandlers{
		repo: repo,
	}
}

func (o *OrderCommandHandlers) Create(c echo.Context) error {
	return o.repo.Create(c)
}
