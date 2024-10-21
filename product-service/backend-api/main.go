package main

import (
	"fmt"
	"net/http"
	"product_service/backend-api/controllers"
	"product_service/backend-api/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("pass1234")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func validateTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		// Cek apakah Authorization header berisi token dengan prefix "Bearer "
		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is missing or malformed"})
			c.Abort()
			return
		}
		tokenString := authHeader[7:]

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Memastikan algoritma yang digunakan adalah HMAC
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

		c.Set("username", claims.Username)
		c.Next()
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Menambahkan header CORS
		c.Header("Access-Control-Allow-Origin", "*") // Ganti '*' dengan domain spesifik jika perlu
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent) // Menangani preflight request
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()

	// Middleware CORS
	router.Use(corsMiddleware())

	// Atur static file untuk gambar produk
	router.Static("/gambarproduk", "./gambarproduk")

	// Hubungkan ke database
	models.ConnectDatabase()

	// Grup route dengan middleware JWT
	protected := router.Group("/api/product")
	protected.Use(validateTokenMiddleware())
	{
		protected.GET("", controllers.FindProduct)
		protected.POST("", controllers.AddProduct)
		protected.PUT("/:id", controllers.EditProduct)
	}

	// Jalankan server pada port 3000
	router.Run(":3000")
}
