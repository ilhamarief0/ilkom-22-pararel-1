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
	// Koneksi ke database untuk migrasi SQL manual
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/ecommerce")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Jalankan migrasi manual
	migrations.Migrate(db)

	// Koneksi database untuk GORM
	models.ConnectDatabase()

	// Inisialisasi router Gin
	r := gin.Default()
	r.Static("/api/gambarproduk", "./gambarproduk")

	// Inisialisasi controller
	productController := &controllers.ProductController{
		DB: models.DB, // Menggunakan koneksi GORM
	}

	// Routes
	r.GET("/api/products", productController.GetAllProducts)
	r.POST("/api/products", productController.AddProduct)
	r.PUT("/api/products/:id", productController.EditProduct)

	// Menjalankan server pada port 3010
	r.Run(":3010")
}
