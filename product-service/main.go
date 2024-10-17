package main

import (
	"product_service/backend-api/controllers"
	"product_service/backend-api/models"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()
	models.ConnectDatabase()

	router.GET("/api/posts", controllers.Finpost)

	router.Run(":3000")
}