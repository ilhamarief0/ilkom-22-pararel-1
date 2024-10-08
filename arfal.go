package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Shipping struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	OrderID uint   `json:"order_id"`
	Address string `json:"address"`
	Status  string `json:"status"`
}

func main() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/ecommerce?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&Shipping{})

	router := gin.Default()
	router.POST("/shipping", createShipping)
	router.Run(":8084") // Shipping service on port 8084
}

func createShipping(c *gin.Context) {
	var shipping Shipping
	if err := c.ShouldBindJSON(&shipping); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&shipping)
	c.JSON(http.StatusCreated, shipping)
}