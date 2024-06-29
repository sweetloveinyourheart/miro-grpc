package utils

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	configs "github.com/sweetloveinyourheart/miro-whiteboard/common/configs"
)

// Validate JWT tokens
func ValidateToken(tokenString string) (*jwt.Token, error) {
	jwtSecret := configs.GetAuthConfig().JwtSecret

	// Get token from Bearer format
	isBearerToken := strings.HasPrefix(tokenString, "Bearer")
	if !isBearerToken {
		return nil, errors.New("invalid token")
	}

	jwtToken, found := strings.CutPrefix(tokenString, "Bearer ")
	if !found {
		return nil, errors.New("invalid token")
	}

	// Parse the token
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	// Validate the token and check the claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if exp < float64(time.Now().Unix()) {
				return nil, errors.New("token is expired")
			}
		}
		return token, nil
	}

	return nil, errors.New("invalid token")
}
