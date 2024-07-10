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
	newUser := services.User{
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

func (s *Server) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	user := services.UserCredential{
		Email:    in.Email,
		Password: in.Password,
	}

	credentials, err := s.svc.GetAuthCredentials(user)
	if err != nil {
		return &pb.SignInResponse{}, err
	}

	response := pb.SignInResponse{
		AccessToken:  credentials.AccessToken,
		RefreshToken: credentials.RefreshToken,
	}

	return &response, nil
}

func (s *Server) GetProfile(ctx context.Context, in *pb.GetProfileRequest) (*pb.ProfileResponse, error) {
	user, err := s.svc.GetUserInfo(in.UserId)
	if err != nil {
		return &pb.ProfileResponse{}, err
	}

	response := pb.ProfileResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	return &response, nil
}
