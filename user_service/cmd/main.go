package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
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

	connStr := os.Getenv("POSTGRESQL_CONN")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := grpc.NewServer()
	svc := server.CreateUserServer(db)
	pb.RegisterUserServiceServer(s, svc)

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
