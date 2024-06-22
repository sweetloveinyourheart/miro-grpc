package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sweetloveinyourheart/miro-whiteboard/gateway/internal/clients"
	"github.com/sweetloveinyourheart/miro-whiteboard/gateway/internal/routers"
)

func main() {
	app := fiber.New()

	v1Routers := app.Group("v1")

	// Create gRPC clients
	userServiceClient, userServiceConn, err := clients.NewUserServiceClient()
	defer userServiceConn.Close()

	if err != nil {
		log.Fatalf("Failed to establish connection with other service, error: %s", err.Error())
	}

	// Create routers
	routers.CreateUserRouters(v1Routers, userServiceClient)

	app.Listen(":9000")
}
