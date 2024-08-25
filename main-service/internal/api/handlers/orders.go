package handlers

import (
	"fmt"
	"net/http"

	"order-service/commands"
	"order-service/models"
	"order-service/pkg/kafka"
	"order-service/pkg/nats"
	"order-service/pkg/rabbitmq"
	"order-service/pkg/rest"
	"order-service/repositories"
	"order-service/repositories/interfaces"

	"github.com/labstack/echo/v4"
)

func CreateOrder(c echo.Context) error {
	var req = new(models.OrderRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	var messageRepo interfaces.MessageInterfaces
	switch req.SendBy {
	case "nats":
		// Get NATS connection
		nc, err := nats.ConnectToNATS()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		defer nc.Close()

		messageRepo = repositories.NewNATSRepository(nc)
	case "kafka":
		// Get Kafka connection
		producer, err := kafka.ConnectToKafka()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		defer producer.Close()

		messageRepo = repositories.NewKafkaRepository(producer)
	case "rest":
		var client = rest.CreateNewRestClient()
		messageRepo = repositories.NewRestRepository(client)
	case "rabbitmq":
		// Get Rabbit MQ connection
		conn, err := rabbitmq.ConnectToRabbitMQ()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		defer conn.Close()

		messageRepo = repositories.NewRabbitMQRepository(conn)
	default:
		return c.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("invalid send type: %s", req.SendBy)})
	}

	var natsCommandHandlers = commands.NewMessageCommandHandlers(messageRepo)

	var orderRepo = repositories.NewOrderRepository(natsCommandHandlers)
	var orderCommandHandler = commands.NewOrderCommandHandlers(orderRepo)

	// Create order
	if err := orderCommandHandler.Create(req.OrderID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not create order"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "order created successfully"})
}
