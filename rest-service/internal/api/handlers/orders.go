package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderMessage struct {
	OrderID string `json:"order_id"`
}

func CreateOrder(c echo.Context) error {
	var req = new(OrderMessage)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	handleMessage(req.OrderID)

	return c.JSON(http.StatusOK, map[string]string{"order created": req.OrderID})
}

func handleMessage(data string) {
	// Process the message
	log.Printf("Processing message: %s", data)
}
