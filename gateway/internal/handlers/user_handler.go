package handlers

import (
	"github.com/gofiber/fiber/v2"
	pb "github.com/sweetloveinyourheart/miro-whiteboard/common/api"
)

type UserHandler struct {
	c pb.UserServiceClient
}

type IUserHandler interface {
	Register(c *fiber.Ctx) error
}

func NewUserHandler() IUserHandler {
	return &UserHandler{}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	return &fiber.Error{}
}
