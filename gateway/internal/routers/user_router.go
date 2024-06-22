package routers

import (
	"github.com/gofiber/fiber/v2"
	pb "github.com/sweetloveinyourheart/miro-whiteboard/common/api"
	"github.com/sweetloveinyourheart/miro-whiteboard/gateway/internal/handlers"
)

func CreateUserRouters(r fiber.Router, c pb.UserServiceClient) {
	routes := r.Group("/users")
	userHandler := handlers.NewUserHandler()

	routes.Get("/register", userHandler.Register)
}
