package main

import (
	"net/http"
	"product_service/backend-api/controllers"
	"product_service/backend-api/models"

	"github.com/gin-gonic/gin"
)

// Removed JWT middleware and related code

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(corsMiddleware())

	router.Static("/gambarproduk", "./gambarproduk")
	models.ConnectDatabase()

	// No JWT validation here, assume the Gateway does that
	productRoutes := router.Group("/api/product")
	{
		productRoutes.GET("", controllers.FindProduct)
		productRoutes.POST("", controllers.AddProduct)
		productRoutes.PUT("/:id", controllers.EditProduct)
	}

	router.Run(":3010")
}
