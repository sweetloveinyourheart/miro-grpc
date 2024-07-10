package clients

import (
	"flag"
	"fmt"

	pb "github.com/sweetloveinyourheart/miro-whiteboard/common/api"
	configs "github.com/sweetloveinyourheart/miro-whiteboard/common/configs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserServiceClient() (pb.UserServiceClient, *grpc.ClientConn, error) {
	// Set up a connection to the server.
	serverPost := configs.GetGRPCConfig().UserServicePort
	addr := flag.String("user_service_addr", fmt.Sprintf("localhost:%d", serverPost), "the address to connect to")

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	return pb.NewUserServiceClient(conn), conn, err
}
