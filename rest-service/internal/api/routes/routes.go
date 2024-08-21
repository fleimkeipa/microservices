package routes

import (
	"rest-service/internal/api/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoutes(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Protected routes
	var r = e.Group("/rest/orders")

	r.POST("", handlers.CreateOrder)
}
