package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"

	pb "github.com/sweetloveinyourheart/miro-whiteboard/common/api"
	configs "github.com/sweetloveinyourheart/miro-whiteboard/common/configs"
	"github.com/sweetloveinyourheart/miro-whiteboard/user_service/internal/server"
)

func main() {
	serverPost := configs.GetGRPCConfig().UserServicePort

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPost))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server.Server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
