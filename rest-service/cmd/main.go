package main

import (
	"rest-service/internal/api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	var e = echo.New()
	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":8081"))
}
