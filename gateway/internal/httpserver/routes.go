package httpserver

import (
	"github.com/labstack/echo/v4"
	"github.com/shiiyan/my-simple-api-platform/gateway/internal/apis/external/handlers"
)

func RegisterRoutes(e *echo.Echo, summaryHandler *handlers.SummaryHandler) {
	e.GET("/api/summary/:id", summaryHandler.GetSummary)
}
