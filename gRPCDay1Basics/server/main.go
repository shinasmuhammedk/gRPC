package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpcBasics/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	fmt.Println("Request received:", req.GetName())

	return &pb.UserResponse{
		Message: "Hello " + req.GetName(),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{})

	fmt.Println("🚀 Server running on port 50051")
	grpcServer.Serve(lis)
}