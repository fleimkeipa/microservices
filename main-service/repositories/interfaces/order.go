package interfaces

import "github.com/labstack/echo/v4"

type OrderInterfaces interface {
	Create(c echo.Context) error
}
