package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // Variabel global untuk koneksi database
type ProductWithUser struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Price    int    `json:"price"`
	Stock    int    `json:"stock"`
	Image    string `json:"image"`
	Username string `json:"owner_username"`
}

// Product model untuk tabel produk
type Product struct {
	ID      int    `gorm:"primaryKey"`
	Title   string `gorm:"type:varchar(255);not null"`
	Content string `gorm:"type:varchar(255);not null"`
	Image   string `gorm:"type:varchar(255);not null"`
	Price   int    `gorm:"not null"`
	Stock   int    `gorm:"not null"`
	UserID  int    `gorm:"not null"`
}

// ConnectDatabase menginisialisasi koneksi ke database
func ConnectDatabase() {
	dsn := "root:@tcp(localhost:3306)/ecommerce?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = database
	log.Println("Database connection and migration completed successfully.")
}
