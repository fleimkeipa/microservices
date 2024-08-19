package main

import (
	"order-service/internal/api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
