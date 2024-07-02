package server

import (
	pb "github.com/sweetloveinyourheart/miro-whiteboard/common/api"
)

type Server struct {
	pb.UnimplementedBoardServiceServer
}

func CreateUserServer() *Server {
	return &Server{}
}
