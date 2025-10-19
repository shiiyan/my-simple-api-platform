package httpserver

import (
	"github.com/labstack/echo/v4"
	"github.com/shiiyan/my-simple-api-platform/service1/internal/port/http/handler"
)

func RegisterRoutes(e *echo.Echo, userHandler *handler.UserHandler) {
	e.GET("/users/:id", userHandler.GetUser)
}
