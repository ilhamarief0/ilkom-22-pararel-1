package handlers

import (
	"context"
	"log"
	"net/http"

	pb "auth-service/proto"
	"auth-service/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	UserServiceClient pb.UserServiceClient
}

func NewAuthHandler(client pb.UserServiceClient) *AuthHandler {
	return &AuthHandler{
		UserServiceClient: client,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginData struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	// Bind form-data
	if err := c.ShouldBind(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	req := &pb.GetUserByUsernameRequest{Username: loginData.Username}
	resp, err := h.UserServiceClient.GetUserByUsername(context.Background(), req)
	if err != nil {
		log.Printf("Error getting user from user-service: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	user := resp.User // Access the User field from the response
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateJWT(loginData.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
