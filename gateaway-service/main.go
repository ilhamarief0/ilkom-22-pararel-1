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
	r.GET("/users", authMiddleware, userHandler.ListUsers)
	r.GET("/users/:id", authMiddleware, userHandler.GetUser)
	r.POST("/users", authMiddleware, userHandler.CreateUser)

	// Proxy POST requests to the product service
	productServiceURL := "http://localhost:3010"
	r.POST("/api/product", func(c *gin.Context) {
		proxyToService(c, productServiceURL, "/api/product")
	})

	// Proxy POST requests to the auth service
	authServiceURL := "http://localhost:9000"
	r.POST("/api/login", func(c *gin.Context) {
		proxyToService(c, authServiceURL, "/api/login")
	})

	// Run REST API on port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func proxyToService(c *gin.Context, target, path string) {
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Printf("Error parsing target URL: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid target URL"})
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// Capture errors from the proxy
	proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
		log.Printf("Error during proxy request to %s: %v", target, err)
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to proxy request", "details": err.Error()})
	}

	// Modify the request URL path to match the target service
	c.Request.URL.Path = path
	proxy.ServeHTTP(c.Writer, c.Request)
}
