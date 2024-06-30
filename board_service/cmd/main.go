package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sweetloveinyourheart/miro-whiteboard/board_service/db"
	"github.com/sweetloveinyourheart/miro-whiteboard/common/configs"
	"google.golang.org/grpc"
)

func main() {
	mongoClient := db.InitMongoConnection()
	defer mongoClient.Disconnect(context.TODO())

	serverPost := configs.GetGRPCConfig().UserServicePort

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPost))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
