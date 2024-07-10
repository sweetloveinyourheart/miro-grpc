package handlers

import (
	pb "github.com/sweetloveinyourheart/miro-whiteboard/common/api"
)

type BoardHandler struct {
	c pb.BoardServiceClient
}

type IBoardHandler interface {
}

func NewBoardHandler(client *pb.BoardServiceClient) IBoardHandler {
	return &BoardHandler{
		c: *client,
	}
}
