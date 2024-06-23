package server

import (
	"context"

	pb "github.com/sweetloveinyourheart/miro-whiteboard/common/api"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func (s *Server) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	response := pb.RegisterResponse{
		Success: true,
		Message: "register successfully !",
	}

	return &response, nil
}
