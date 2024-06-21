package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/joho/godotenv/autoload"
	pb "github.com/sweetloveinyourheart/miro-whiteboard/common/api"
	"github.com/sweetloveinyourheart/miro-whiteboard/user_service/internal/server"
	"google.golang.org/grpc"
)

func main() {
	serverPost := os.Getenv("SERVER_PORT")

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", serverPost))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server.RPCServer{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
