package handlers

import (
	"context"
	"net/http"
	"strconv"

	pb "gateway-service/proto"

	"github.com/gin-gonic/gin"
)

type GatewayHandler struct {
	UserService pb.UserServiceClient
}

func (h *GatewayHandler) GetUser(c *gin.Context) {
	// Mendapatkan ID dari parameter URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam) // Konversi ID dari string ke integer
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Memanggil gRPC GetUser
	req := &pb.UserRequest{Id: int32(id)}
	res, err := h.UserService.GetUser(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	// Membuat response JSON
	c.JSON(http.StatusOK, res.User) // Return user dengan role yang sudah disertakan
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, gin.H{"success": res.Success, "message": res.Message})
}
