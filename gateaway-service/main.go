package main

import (
	"gateway-service/config"
	"gateway-service/handlers"
	"gateway-service/proto"

	"github.com/gin-gonic/gin"
)

func main() {
	// Koneksi ke gRPC user-service
	grpcConn := config.ConnectToGRPCServer("localhost:50051")
	defer grpcConn.Close()

	// Inisialisasi gRPC client
	userServiceClient := proto.NewUserServiceClient(grpcConn)

	// Inisialisasi Gin
	r := gin.Default()

	// Register handler
	userHandler := &handlers.UserHandler{UserService: userServiceClient}
	r.GET("/users/:id", userHandler.GetUser)

	// Jalankan REST API
	r.Run(":8080") // REST API berjalan di port 8080
}
