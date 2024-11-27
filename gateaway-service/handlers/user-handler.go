package handlers

import (
	"context"
	"gateway-service/proto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService proto.UserServiceClient
}

func (h *UserHandler) GetUser(c *gin.Context) {
	// Ambil parameter `id` dari URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Panggil gRPC GetUser
	res, err := h.UserService.GetUser(context.Background(), &proto.UserRequest{Id: int32(id)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Kembalikan response ke klien
	c.JSON(http.StatusOK, gin.H{
		"id":       res.User.Id,
		"username": res.User.Username,
		"email":    res.User.Email,
		"role":     res.User.Role,
	})
}
