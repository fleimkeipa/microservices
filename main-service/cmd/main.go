package main

import (
	"order-service/internal/api/routes"
	"order-service/pkg/nats"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	routes.SetupRoutes(e)

	nc := nats.ConnectToNATS()
	defer nc.Close()

	e.Logger.Fatal(e.Start(":8080"))
}
