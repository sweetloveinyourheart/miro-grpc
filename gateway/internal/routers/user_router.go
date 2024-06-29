package routers

import (
	"github.com/gofiber/fiber/v2"
	pb "github.com/sweetloveinyourheart/miro-whiteboard/common/api"
	"github.com/sweetloveinyourheart/miro-whiteboard/gateway/internal/handlers"
	"github.com/sweetloveinyourheart/miro-whiteboard/gateway/internal/middlewares"
)

func CreateUserRouters(r fiber.Router, c *pb.UserServiceClient) {
	routes := r.Group("/users")
	userHandler := handlers.NewUserHandler(c)

	routes.Post("/register", userHandler.Register)
	routes.Post("/sign-in", userHandler.SignIn)
	routes.Get("/profile", middlewares.AuthGuard, userHandler.GetProfile)
}
