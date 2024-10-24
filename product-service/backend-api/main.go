package main

import (
	"product_service/backend-api/controllers"
	"product_service/backend-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/api/gambarproduk", "./gambarproduk")
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
