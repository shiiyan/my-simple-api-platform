package httpserver

import "github.com/labstack/echo/v4"

func NewServer() *echo.Echo {
	return echo.New()
}
