package main

import (
	"github.com/shiiyan/my-simple-api-platform/gateway/internal/apis/external/handlers"
	"github.com/shiiyan/my-simple-api-platform/gateway/internal/common/clients"
	"github.com/shiiyan/my-simple-api-platform/gateway/internal/httpserver"
)

func main() {
	e := httpserver.NewServer()

	service1Client := clients.NewService1Client("http://localhost:8080") // adjust port as needed
	service2Client := clients.NewService2Client("http://localhost:8081") // adjust port as needed

	summaryHandler := handlers.NewSummaryHandler(service1Client, service2Client)

	httpserver.RegisterRoutes(e, summaryHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
