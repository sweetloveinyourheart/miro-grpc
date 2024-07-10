package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sweetloveinyourheart/miro-whiteboard/gateway/internal/clients"
	"github.com/sweetloveinyourheart/miro-whiteboard/gateway/internal/routers"
)

func main() {
	app := fiber.New()

	v1Routers := app.Group("/api/v1")

	// Create gRPC clients
	userServiceClient, userServiceConn, userServiceErr := clients.NewUserServiceClient()
	defer userServiceConn.Close()
	if userServiceErr != nil {
		log.Fatalf("Failed to establish connection with user service, error: %s", userServiceErr.Error())
	}

	boardServiceClient, boardServiceConn, boardServiceErr := clients.NewBoardServiceClient()
	defer boardServiceConn.Close()
	if boardServiceErr != nil {
		log.Fatalf("Failed to establish connection with board service, error: %s", userServiceErr.Error())
	}

	// Create routers
	routers.CreateUserRouters(v1Routers, &userServiceClient)
	routers.CreateBoardRouters(v1Routers, &boardServiceClient)

	app.Listen(":9000")
}
