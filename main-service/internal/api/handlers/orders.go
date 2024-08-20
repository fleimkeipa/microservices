package handlers

import (
	"net/http"

	"order-service/commands"
	"order-service/pkg/nats"
	"order-service/repositories"

	"github.com/labstack/echo/v4"
)

type OrderRequest struct {
	OrderID string `json:"order_id"`
}

func CreateOrder(c echo.Context) error {
	// Get NATS connection
	nc := nats.ConnectToNATS()
	defer nc.Close()

	var orderRepo = repositories.NewOrderRepository()
	var orderCommandHandler = commands.NewOrderCommandHandlers(orderRepo)

	// Create order
	if err := orderCommandHandler.Create(c); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not create order"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "order created successfully"})
}
