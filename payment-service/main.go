package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Payment struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	UserID uint    `json:"user_id"`
	Amount float64 `json:"amount"`
}

func main() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/ecommerce?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&Payment{})

	router := gin.Default()
	router.POST("/payments", createPayment)
	router.Run(":8083") // Payment service on port 8083
}

func createPayment(c *gin.Context) {
	var payment Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&payment)
	c.JSON(http.StatusCreated, payment)
}
