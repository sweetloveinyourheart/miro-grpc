package utils

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"
)

func GetMetadata(ctx context.Context) (metadata.MD, error) {
	// Retrieve metadata from the incoming context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	return md, nil
}

func GetAuthorizedUser(ctx context.Context) (string, error) {
	email := ""

	// Get context metadata
	md, err := GetMetadata(ctx)
	if err != nil {
		return email, err
	}

	// Extract the email from metadata
	if val, ok := md["email"]; ok && len(val) > 0 {
		email = val[0]
	} else {
		return email, fmt.Errorf("email not provided in metadata")
	}

	return email, nil
}
