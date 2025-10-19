package main

import (
	"github.com/shiiyan/my-simple-api-platform/service1/internal/application"
	"github.com/shiiyan/my-simple-api-platform/service1/internal/port/db"
	"github.com/shiiyan/my-simple-api-platform/service1/internal/port/http/config"
	"github.com/shiiyan/my-simple-api-platform/service1/internal/port/http/handler"
	"github.com/shiiyan/my-simple-api-platform/service1/internal/port/http/httpserver"
)

func main() {
	e := httpserver.NewServer()

	userRepo := db.NewInMemoryUserRepo()
	userUseCase := application.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)

	httpserver.RegisterRoutes(e, userHandler)

	cfg := config.NewHTTPConfig()
	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
