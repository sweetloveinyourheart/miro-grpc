package server

import (
	pb "github.com/sweetloveinyourheart/miro-whiteboard/common/api"
)

type RPCServer struct {
	pb.UnimplementedUserServiceServer
}
