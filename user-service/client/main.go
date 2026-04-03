package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "user-service/proto"

	"google.golang.org/grpc"
)

func main() {
	// Connect to server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	// Timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call CreateUser
	res, err := client.CreateUser(ctx, &pb.User{
		Name: "Shinas",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response:", res)
}