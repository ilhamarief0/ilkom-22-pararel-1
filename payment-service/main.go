package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors" // Using gin-contrib/cors
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Payment struct {
	ProductName string  `json:"productName" binding:"required"`
	UserID      int     `json:"userId" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required"`
	TotalPrice  float64 `json:"totalPrice" binding:"required"`
	Status      string  `json:"status"`
}

var db *sql.DB

func main() {
	var err error
	// Connect to MySQL database
	dsn := "root:@tcp(127.0.0.1:3306)/ecommerce"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize Gin router
	r := gin.Default()

	// Configure CORS middleware using gin-contrib/cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Allow your frontend origin
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	// Route to handle payment service (order creation)
	r.POST("/payments", createPayment)

	// Start the server
	r.Run(":8085")
}

// createPayment handles new payments/orders
func createPayment(c *gin.Context) {
	var newPayment Payment

	// Parse JSON request body into the Payment struct
	if err := c.ShouldBindJSON(&newPayment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment data"})
		return
	}

	// Set default status if not provided
	if newPayment.Status == "" {
		newPayment.Status = "waiting for payment"
	}

	// Insert the payment data into the database
	query := `INSERT INTO payments (product_name, user_id, quantity, total_price, status, created_at)
		VALUES (?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, newPayment.ProductName, newPayment.UserID, newPayment.Quantity, newPayment.TotalPrice, newPayment.Status, time.Now())

	if err != nil {
		log.Println("Error inserting payment:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment"})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{
		"message":     "Payment created successfully",
		"productName": newPayment.ProductName,
		"userId":      newPayment.UserID,
		"quantity":    newPayment.Quantity,
		"totalPrice":  newPayment.TotalPrice,
		"status":      newPayment.Status,
	})
}
