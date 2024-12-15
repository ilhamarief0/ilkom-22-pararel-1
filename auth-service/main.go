package main

import (
	"log"

	"auth-service/handlers"
	pb "auth-service/proto" // Ensure this path matches the go_package

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	// Connect to the user-service gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to user-service: %v", err)
	}
	defer conn.Close()

	// Create a new UserService client
	userServiceClient := pb.NewUserServiceClient(conn)

	// Initialize the AuthHandler with the gRPC client
	authHandler := handlers.NewAuthHandler(userServiceClient)

	// Set up the Gin router
	r := gin.Default()
	r.POST("/api/login", authHandler.Login)

	// Start the HTTP server
	if err := r.Run(":9000"); err != nil {
		log.Fatalf("Failed to run auth-service: %v", err)
	}
}
