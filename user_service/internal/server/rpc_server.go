package server

import (
	"context"
	"database/sql"

	pb "github.com/sweetloveinyourheart/miro-whiteboard/common/api"
	"github.com/sweetloveinyourheart/miro-whiteboard/user_service/internal/services"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	svc services.IUserServices
}

func CreateUserServer(db *sql.DB) *Server {
	return &Server{
		svc: services.CreateUserService(db),
	}
}

func (s *Server) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	newUser := services.NewUser{
		FirstName: in.Email,
		LastName:  in.LastName,
		Email:     in.Email,
		Password:  in.Password,
	}

	success, err := s.svc.CreateNewUser(newUser)
	if err != nil {
		response := pb.RegisterResponse{
			Success: false,
			Message: "register failed !",
		}

		return &response, nil
	}

	if !success {
		response := pb.RegisterResponse{
			Success: false,
			Message: "user already exist !",
		}

		return &response, nil
	}

	response := pb.RegisterResponse{
		Success: true,
		Message: "register successfully !",
	}

	return &response, nil
}
