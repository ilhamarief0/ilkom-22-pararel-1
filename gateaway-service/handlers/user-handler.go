package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	pb "gateway-service/proto"

	"github.com/gin-gonic/gin"
)

type GatewayHandler struct {
	UserService pb.UserServiceClient
}

func (h *GatewayHandler) ListUsers(c *gin.Context) {
	req := &pb.ListUsersRequest{}
	res, err := h.UserService.ListUsers(context.Background(), req)
	if err != nil {
		log.Printf("Error calling ListUsers gRPC: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list users"})
		return
	}

	c.JSON(http.StatusOK, res.Users)
}

func (h *GatewayHandler) GetUser(c *gin.Context) {
	// Get ID from URL parameter
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Call gRPC GetUser
	req := &pb.UserRequest{Id: int32(id)}
	res, err := h.UserService.GetUser(context.Background(), req)
	if err != nil {
		log.Printf("Error calling GetUser: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	// Return user as JSON
	c.JSON(http.StatusOK, res.User)
}

func (h *GatewayHandler) CreateUser(c *gin.Context) {
	// Bind JSON request to CreateUserRequest
	var req pb.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call gRPC CreateUser
	res, err := h.UserService.CreateUser(context.Background(), &req)
	if err != nil {
		log.Printf("Error calling CreateUser: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, gin.H{"success": res.Success, "message": res.Message})
}
