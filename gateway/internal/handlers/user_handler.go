package handlers

import (
	"context"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	pb "github.com/sweetloveinyourheart/miro-whiteboard/common/api"
	"github.com/sweetloveinyourheart/miro-whiteboard/gateway/internal/requests"
	"github.com/sweetloveinyourheart/miro-whiteboard/gateway/internal/responses"
	"github.com/sweetloveinyourheart/miro-whiteboard/gateway/internal/utils"
)

type UserHandler struct {
	c pb.UserServiceClient
}

type IUserHandler interface {
	Register(c *fiber.Ctx) error
}

func NewUserHandler(client *pb.UserServiceClient) IUserHandler {
	return &UserHandler{
		c: *client,
	}
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	// Get request body
	var newUser requests.RegisterRequestData
	unMarshalErr := json.Unmarshal(ctx.Body(), &newUser)
	if unMarshalErr != nil {
		return ctx.Status(400).JSON(responses.RegisterResponseData{
			Success: false,
			Message: fiber.ErrBadRequest.Error(),
		})
	}

	// Validate request body
	if errs := utils.Validate(newUser); len(errs) > 0 && errs[0].Error {
		validationMessage := utils.CreateValidationMessage(errs)
		return ctx.Status(400).JSON(responses.RegisterResponseData{
			Success: false,
			Message: validationMessage,
		})
	}

	grpcData := pb.RegisterRequest{
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
		Password:  newUser.Password,
	}
	grpcContext := context.Background()
	result, grpcErr := h.c.Register(grpcContext, &grpcData)
	if grpcErr != nil {
		return fiber.ErrInternalServerError
	}

	response := responses.RegisterResponseData{
		Success: result.Success,
		Message: result.Message,
	}
	return ctx.Status(201).JSON(response)
}
