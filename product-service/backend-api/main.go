package main

import (
	"fmt"
	"net/http"
	"product_service/backend-api/controllers"
	"product_service/backend-api/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("pass1234")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func validateTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is missing or malformed"})
			c.Abort()
			return
		}

		// Hanya mengambil bagian token tanpa prefix "Bearer "
		tokenString := authHeader[7:]

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Pastikan algoritma yang digunakan adalah HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Jika valid, lanjut ke handler berikutnya
		c.Set("username", claims.Username)
		c.Next()
	}
}

func main() {
	router := gin.Default()
	models.ConnectDatabase()

	// Menggunakan middleware JWT untuk route /api/product
	protected := router.Group("/api/product")
	protected.Use(validateTokenMiddleware())
	{
		protected.GET("", controllers.Findpost)
		protected.POST("", controllers.AddPost)
		protected.PUT("", controllers.EditPost)
	}

	router.Run(":3000")
}
 