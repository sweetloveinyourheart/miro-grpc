package server

import (
	"context"
	"fmt"

	"github.com/sweetloveinyourheart/miro-whiteboard/board_service/internal/services"
	"github.com/sweetloveinyourheart/miro-whiteboard/board_service/internal/utils"
	pb "github.com/sweetloveinyourheart/miro-whiteboard/common/api"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	pb.UnimplementedBoardServiceServer
	svc services.IBoardService
}

func CreateUserServer(client *mongo.Client) *Server {
	return &Server{
		svc: services.NewBoardService(client),
	}
}

func (s *Server) CreateBoard(ctx context.Context, in *pb.CreateBoardRequest) (*pb.CreateBoardResponse, error) {
	user, err := utils.GetAuthorizedUser(ctx)
	if err != nil {
		return nil, fmt.Errorf("unauthorized")
	}

	board := services.BoardInfo{
		Title:       in.Title,
		Description: in.Description,
		CreatedBy:   user,
	}

	success, err := s.svc.CreateBoard(board)
	if err != nil {
		return nil, err
	}

	response := pb.CreateBoardResponse{
		Success: success,
		Message: "New board created",
	}

	return &response, nil
}

func GetBoard() {}

func DeleteBoard() {}
