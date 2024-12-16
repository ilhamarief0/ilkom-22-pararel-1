package main

import (
	"gateway-service/config"
	"gateway-service/handlers"
	user "gateway-service/proto"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("jWtTokenKel1")

func main() {
	// Connect to gRPC user-service
	grpcConn := config.ConnectToGRPCServer("localhost:50051")
	defer grpcConn.Close()

	// Initialize gRPC client
	userServiceClient := user.NewUserServiceClient(grpcConn)

	// Initialize Gin
	r := gin.Default()

	// JWT Middleware
	authMiddleware := func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims := &jwt.RegisteredClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("username", claims.Subject)
		c.Next()
	}

	// Register user handlers with JWT middleware
	userHandler := &handlers.GatewayHandler{UserService: userServiceClient}
	r.GET("/api/users", authMiddleware, userHandler.ListUsers)
	r.GET("/api/users/:id", authMiddleware, userHandler.GetUser)
	r.POST("/api/users", authMiddleware, userHandler.CreateUser)
	r.PUT("/api/users/:id", authMiddleware, userHandler.UpdateUser)
	r.DELETE("/api/users/:id", authMiddleware, userHandler.DeleteUser)

	// Proxy POST requests to the product service
	productServiceURL := "http://localhost:3010"
	r.POST("/api/products", func(c *gin.Context) {
		proxyToService(c, productServiceURL, "/api/products")
	})
	r.GET("/api/products", authMiddleware, func(c *gin.Context) {
		proxyToService(c, productServiceURL, "/api/products")
	})
	r.GET("/api/products/:id", func(c *gin.Context) {
		proxyToService(c, productServiceURL, "/api/products/:id")
	})
	r.DELETE("/api/products/:id", authMiddleware, func(c *gin.Context) {
		proxyToService(c, productServiceURL, "/api/products/:id")
	})
	r.PUT("/api/products/:id", authMiddleware, func(c *gin.Context) {
		proxyToService(c, productServiceURL, "/api/products/:id")
	})

	// Proxy POST requests to the auth service
	authServiceURL := "http://localhost:9000"
	r.POST("/api/login", func(c *gin.Context) {
		proxyToService(c, authServiceURL, "/api/login")
	})
	// Proxy POST requests to the auth service
	paymentServiceURL := "http://localhost:9010"
	r.POST("/api/payments", func(c *gin.Context) {
		proxyToService(c, paymentServiceURL, "/api/payments")
	})

	// Run REST API on port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func proxyToService(c *gin.Context, serviceURL string, path string) {
	// Parse the service URL
	target, err := url.Parse(serviceURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid service URL"})
		return
	}

	// Create a reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(target)

	// Update the request URL to forward to the correct path
	c.Request.URL.Path = path

	// Serve the request using the proxy
	proxy.ServeHTTP(c.Writer, c.Request)
}
