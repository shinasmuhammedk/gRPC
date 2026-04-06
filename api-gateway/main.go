package main

import (
	"api-gateway/handler"
	"api-gateway/middleware"
	pb "api-gateway/proto"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {

	//Connect to userService
	userConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect UserService: ", err)
	}
	userClient := pb.NewUserServiceClient(userConn)

	orderConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect OrderService: ", err)
	}
	orderClient := pb.NewOrderServiceClient(orderConn)

	//Gin Router
	r := gin.Default()
    
	//Register handler
	handler.RegisterUserRoutes(r, userClient)
    
    //Protected routes
    auth := r.Group("/")
    auth.Use(
        middleware.AuthMiddleWare(),
        middleware.RateLimitMiddleware(),
    )
    
	handler.RegisterOrderRoutes(auth, orderClient)

	log.Println("API Gateway running on : 8080")

	r.Run(":8080")
}
