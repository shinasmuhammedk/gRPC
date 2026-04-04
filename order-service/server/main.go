package main

import (
	"log"
	"net"
	"order-service/db"
	pb "order-service/proto"

	userpb "order-service/proto"
	"order-service/service"

	"google.golang.org/grpc"
)

func main() {
	database := db.ConnectDB()

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to UserSevice", err)
	}

	userClient := userpb.NewUserServiceClient(conn)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen:", err)
	}

	grpcServer := grpc.NewServer()

	orderService := &service.OrderService{
		DB:         database,
		UserClient: userClient,
	}

    pb.RegisterOrderServiceServer(grpcServer, orderService)
    
    log.Println("OrderService running on :50052")
    
    if err := grpcServer.Serve(lis); err != nil{
        log.Fatal(err)
    }
}
