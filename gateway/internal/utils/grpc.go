package utils

import (
	"context"

	"google.golang.org/grpc/metadata"
)

func CreateAuthContext(userId string) context.Context {
	// Create metadata with user information.
	md := metadata.New(map[string]string{
		"user_id": userId,
	})

	ctx := metadata.NewOutgoingContext(context.Background(), md)

	return ctx
}
