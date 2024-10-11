package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

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

func createProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result, err := db.Exec("INSERT INTO products (name, description, price, stock) VALUES (?, ?, ?, ?)", product.Name, product.Description, product.Price, product.Stock)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product", "details": err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get product ID", "details": err.Error()})
		return
	}

	product.ID = int(id)
	c.JSON(http.StatusCreated, product)
}

func getProducts(c *gin.Context) {
	rows, err := db.Query("SELECT id, name, description, price, stock FROM products")
	if err != nil {
		log.Println("Error querying products:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	defer rows.Close()

	products := []Product{}
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock); err != nil {
			log.Println("Error scanning product:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan product"})
			return
		}
		products = append(products, product)
	}

	c.JSON(http.StatusOK, products)
}

func getProduct(c *gin.Context) {
	id := c.Param("id")
	var product Product
	err := db.QueryRow("SELECT id, name, description, price, stock FROM products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func updateProduct(c *gin.Context) {
	id := c.Param("id")
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := db.Exec("UPDATE products SET name = ?, description = ?, price = ?, stock = ? WHERE id = ?", product.Name, product.Description, product.Price, product.Stock, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	product.ID, _ = strconv.Atoi(id)
	c.JSON(http.StatusOK, product)
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")
	_, err := db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080", "*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/products", createProduct)
	r.GET("/products", getProducts)
	r.GET("/products/:id", getProduct)
	r.PUT("/products/:id", updateProduct)
	r.DELETE("/products/:id", deleteProduct)

	// Support preflight requests
	r.OPTIONS("/products", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	r.OPTIONS("/products/:id", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	if err := r.Run(":8082"); err != nil { // Product Service on port 8082
		log.Fatal("Failed to run server:", err)
	}
}
