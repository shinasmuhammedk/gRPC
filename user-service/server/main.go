package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	"user-service/db"
	"user-service/service"

	_ "github.com/lib/pq"

	pb "user-service/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
	dbConn *sql.DB
}

func (s *server) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	id, err := service.CreateUserService(s.dbConn, req.Name)
	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:   int32(id),
		Name: req.Name,
	}, nil
}

func main() {
	// Start listener
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	// Connect DB
	dbConn := db.Connect()

	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Register service
	pb.RegisterUserServiceServer(grpcServer, &server{
		dbConn: dbConn,
	})

	fmt.Println("🚀 Server running on :50051")

	// Start server
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
