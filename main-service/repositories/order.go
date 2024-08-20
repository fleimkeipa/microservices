package repositories

import (
	"log"
	"net/http"

	"order-service/models"
	"order-service/pkg/nats"

	"github.com/labstack/echo/v4"
)

type OrderRepository struct {
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

func (o *OrderRepository) Create(c echo.Context) error {
	var req = new(models.OrderRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	// Get NATS connection
	var nc = nats.ConnectToNATS()
	defer nc.Close()

	// Create order
	var err = nc.Publish("order.created", []byte(req.OrderID))
	if err != nil {
		log.Printf("Failed to publish order.created: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not create order"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "order created successfully"})
}
