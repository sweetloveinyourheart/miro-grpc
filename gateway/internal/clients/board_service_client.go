package clients

import (
	"flag"
	"fmt"

	pb "github.com/sweetloveinyourheart/miro-whiteboard/common/api"
	configs "github.com/sweetloveinyourheart/miro-whiteboard/common/configs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewBoardServiceClient() (pb.BoardServiceClient, *grpc.ClientConn, error) {
	// Set up a connection to the server.
	serverPost := configs.GetGRPCConfig().BoardServicePort
	addr := flag.String("addr", fmt.Sprintf("localhost:%d", serverPost), "the address to connect to")

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	return pb.NewBoardServiceClient(conn), conn, err
}
