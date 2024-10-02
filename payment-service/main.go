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

	// Allow CORS
	r.Use(cors.Default())

	// Routes
	r.POST("/payments", createPayment)
	r.GET("/payments/:userId", getPaymentsByUser) // New route to get payments by user ID

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

func getPaymentsByUser(c *gin.Context) {
	userID := c.Param("userId")

	// Query to select payments based on the userId
	query := `SELECT product_name, quantity, total_price, status FROM payments WHERE user_id = ?`

	rows, err := db.Query(query, userID)
	if err != nil {
		log.Println("Error querying payments:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve payments"})
		return
	}
	defer rows.Close()

	var payments []Payment
	for rows.Next() {
		var payment Payment
		if err := rows.Scan(&payment.ProductName, &payment.Quantity, &payment.TotalPrice, &payment.Status); err != nil {
			log.Println("Error scanning payment:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve payments"})
			return
		}
		payments = append(payments, payment)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error iterating over payment rows:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve payments"})
		return
	}

	// Return payments as JSON
	c.JSON(http.StatusOK, payments)
}
