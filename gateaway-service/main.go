package main

import (
	"gateway-service/config"
	"gateway-service/handlers"
	"gateway-service/proto"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to gRPC user-service
	grpcConn := config.ConnectToGRPCServer("localhost:50051")
	defer grpcConn.Close()

	// Initialize gRPC client
	userServiceClient := proto.NewUserServiceClient(grpcConn)

	// Initialize Gin
	r := gin.Default()

	// Register user handlers
	userHandler := &handlers.GatewayHandler{UserService: userServiceClient}
	r.GET("/users/:id", userHandler.GetUser)
	r.POST("/users", userHandler.CreateUser)

	// Proxy POST requests to the product service
	productServiceURL := "http://localhost:3010"
	r.POST("/api/product", func(c *gin.Context) {
		proxyToProductService(c, productServiceURL)
	})

	// Run REST API on port 8080
	r.Run(":8080")
}

func proxyToProductService(c *gin.Context, target string) {
	targetURL, err := url.Parse(target)
	if err != nil {
		c.JSON(500, gin.H{"error": "Invalid target URL"})
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// Modify the request URL path to match the product service
	c.Request.URL.Path = "/api/product"
	proxy.ServeHTTP(c.Writer, c.Request)
}
