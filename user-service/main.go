package main

import (
	"log"
	"user-service/controllers"
	"user-service/database"
	"user-service/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	database.ConnectDB()

	// Migrate the schema (create the User table)
	database.DB.AutoMigrate(&models.User{})

	// Initialize gin router
	r := gin.Default()

	// Configure CORS with gin-contrib/cors
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3001"}, // Frontend origin
		AllowCredentials: true,
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	}
	r.Use(cors.New(config))

	// Register routes and connect them with controllers
	r.GET("/users", controllers.GetAllUsers)
	r.GET("/user/:id", controllers.GetUser)
	r.POST("/user", controllers.CreateUser)
	r.PUT("/user/:id", controllers.UpdateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)

	// Start the server
	log.Println("Server running on port 8081")
	r.Run(":8081") // Run the gin server
}
