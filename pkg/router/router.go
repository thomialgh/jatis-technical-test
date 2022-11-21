package router

import (
	"jatis/pkg/handler"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	e.GET("/order/:orderID", handler.GetOrderDetail)
}
