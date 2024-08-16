package handlers

import (
	"net/http"

	"order-service/order"
	"order-service/pkg/nats"

	"github.com/labstack/echo/v4"
)

type OrderRequest struct {
	OrderID string `json:"order_id"`
}

func CreateOrder(c echo.Context) error {
	req := new(OrderRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	// Get NATS connection
	nc := nats.ConnectToNATS()
	defer nc.Close()

	// Start OrderService
	orderService := order.NewOrderService(nc)

	// Create order
	err := orderService.CreateOrder(c.Request().Context(), req.OrderID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not create order"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "order created successfully"})
}
