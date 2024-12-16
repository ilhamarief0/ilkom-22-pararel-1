package main

import (
	"database/sql"
	"log"
	"product_service/controllers"
	"product_service/db/migrations"
	"product_service/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connect to the database for manual SQL migrations
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/ecommerce")
	if err != nil {
		log.Fatalf("Failed to connect to database for migrations: %v", err)
	}
	defer db.Close()

	// Run manual migrations
	migrations.Migrate(db)

	// Connect to the database using GORM
	models.ConnectDatabase()

	// Initialize Gin router
	r := gin.Default()
	r.Static("/api/gambarproduk", "./gambarproduk")

	// Initialize product controller with the GORM database connection
	productController := &controllers.ProductController{
		DB: models.DB,
	}

	// Define routes
	r.GET("/api/products", productController.GetAllProducts)
	r.POST("/api/products", productController.AddProduct)
	r.GET("/api/products/:id", productController.GetProductByID)
	r.DELETE("/api/products/:id", productController.DeleteProduct)
	r.PUT("/api/products/:id", productController.EditProduct)

	// Start the server on port 3010
	if err := r.Run(":	"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
