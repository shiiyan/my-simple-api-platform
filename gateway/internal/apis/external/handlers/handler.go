package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shiiyan/my-simple-api-platform/gateway/internal/common/clients"
)

type SummaryHandler struct {
	Service1 *clients.Service1Client
	Service2 *clients.Service2Client
}

func NewSummaryHandler(s1 *clients.Service1Client, s2 *clients.Service2Client) *SummaryHandler {
	return &SummaryHandler{Service1: s1, Service2: s2}
}

func (h *SummaryHandler) GetSummary(c echo.Context) error {
	id := c.Param("id")

	user, err := h.Service1.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "failed to get user"})
	}

	order, err := h.Service2.GetOrderByUserId(user.ID)
	if err != nil {
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "failed to get order"})
	}

	return c.JSON(http.StatusOK, GetSummaryResponse{
		User: User{
			ID:   user.ID,
			Name: user.Name,
		},
		Order: order,
	})
}
