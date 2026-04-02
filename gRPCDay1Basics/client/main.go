package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "grpcBasics/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.UserRequest{
		Name: "Shinas",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response:", res.Message)
}