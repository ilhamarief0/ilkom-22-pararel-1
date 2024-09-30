package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

type Order struct {
	ID         int     `json:"id"`
	UserID     int     `json:"user_id"`
	ProductID  int     `json:"product_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
	OrderDate  string  `json:"order_date"`
	Status     string  `json:"status"`
}

// Initialize a variable to hold your database connection
var db *sql.DB

func init() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/ecommerce" // Update with your MySQL credentials
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Error pinging database:", err)
	}
}

func createOrder(c *gin.Context) {
	var order Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Calculate total price based on the product price and quantity
	var productPrice float64
	err := db.QueryRow("SELECT price FROM products WHERE id = ?", order.ProductID).Scan(&productPrice)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return
	}

	order.TotalPrice = productPrice * float64(order.Quantity)

	result, err := db.Exec("INSERT INTO orders (user_id, product_id, quantity, total_price) VALUES (?, ?, ?, ?)",
		order.UserID, order.ProductID, order.Quantity, order.TotalPrice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get order ID"})
		return
	}

	order.ID = int(id)
	c.JSON(http.StatusCreated, order)
}

func getOrders(c *gin.Context) {
	rows, err := db.Query("SELECT id, user_id, product_id, quantity, total_price, order_date, status FROM orders")
	if err != nil {
		log.Println("Error querying orders:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		if err := rows.Scan(&order.ID, &order.UserID, &order.ProductID, &order.Quantity, &order.TotalPrice, &order.OrderDate, &order.Status); err != nil {
			log.Println("Error scanning order:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan order"})
			return
		}
		orders = append(orders, order)
	}

	c.JSON(http.StatusOK, orders)
}

func main() {
	r := gin.Default()

	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Adjust to your frontend URL
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.POST("/orders", createOrder)
	r.GET("/orders", getOrders)

	if err := r.Run(":8083"); err != nil { // Order Service on port 8082
		log.Fatal("Failed to run server:", err)
	}
}
