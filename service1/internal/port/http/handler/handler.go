package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shiiyan/my-simple-api-platform/service1/internal/application"
)

type UserHandler struct {
	Usecase *application.UserUsecase
}

func NewUserHandler(usecase *application.UserUsecase) *UserHandler {
	return &UserHandler{Usecase: usecase}
}

func (h *UserHandler) GetUser(c echo.Context) error {
	id := c.Param("id")
	user, err := h.Usecase.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}
