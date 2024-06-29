package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/sweetloveinyourheart/miro-whiteboard/gateway/internal/utils"
)

func AuthGuard(ctx *fiber.Ctx) error {
	authorization := ctx.Get("Authorization")

	credentials, err := utils.ValidateToken(authorization)
	if err != nil {
		return ctx.SendStatus(401)
	}

	claims := credentials.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	ctx.Request().Header.Set("user", email)

	return ctx.Next()
}
