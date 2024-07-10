package utils

import (
	"context"
	"fmt"
	"strconv"

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

func GetAuthorizedUser(ctx context.Context) (int32, error) {
	userIdCtx := ""

	// Get context metadata
	md, err := GetMetadata(ctx)
	if err != nil {
		return 0, err
	}

	// Extract the user_id from metadata
	if val, ok := md["user_id"]; ok && len(val) > 0 {
		userIdCtx = val[0]
	} else {
		return 0, fmt.Errorf("email not provided in metadata")
	}

	userId, convertErr := strconv.ParseInt(userIdCtx, 10, 32)
	if convertErr != nil {
		return 0, fmt.Errorf("cannot convert user id")
	}

	return int32(userId), nil
}
