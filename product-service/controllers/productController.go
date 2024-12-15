package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"product_service/models" // Import models package

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ProductController struct {
	DB *gorm.DB
}

type ValidatePostInput struct {
	Title   string `form:"title" json:"title" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
	Price   int    `form:"price" json:"price" binding:"required"`
	Stock   int    `form:"stock" json:"stock" binding:"required"`
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	}
	return "Unknown error"
}

func (pc *ProductController) GetAllProducts(c *gin.Context) {
	// Define the query to fetch product and user information
	query := `
		SELECT p.id, p.title, p.content, p.price, p.stock, p.image, u.username
		FROM product p
		LEFT JOIN users u ON p.user_id = u.id
	`

	// Declare a slice to hold the result inline
	var products []models.ProductWithUser

	// Execute the query and directly scan into the products slice
	if err := pc.DB.Raw(query).Scan(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	// Return the result in JSON format directly
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "List of Products",
		"data":    products,
	})
}

// Handler for adding a new product
func (pc *ProductController) AddProduct(c *gin.Context) {
	var input ValidatePostInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
		return
	}

	imageFolder := "gambarproduk"
	if _, err := os.Stat(imageFolder); os.IsNotExist(err) {
		if err := os.Mkdir(imageFolder, os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create image directory"})
			return
		}
	}

	filename := filepath.Base(file.Filename)
	imagePath := filepath.Join(imageFolder, filename)
	if err := c.SaveUploadedFile(file, imagePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	product := models.Product{
		Title:   input.Title,
		Content: input.Content,
		Price:   input.Price,
		Stock:   input.Stock,
		Image:   filename,
	}

	if err := pc.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Product created successfully",
		"data":    product,
	})
}

// Handler for editing a product
func (pc *ProductController) EditProduct(c *gin.Context) {
	var product models.Product
	if err := pc.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var input ValidatePostInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	file, err := c.FormFile("image")
	if err == nil {
		imageFolder := "gambarproduk"
		filename := filepath.Base(file.Filename)
		imagePath := filepath.Join(imageFolder, filename)

		if err := c.SaveUploadedFile(file, imagePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
			return
		}

		// Delete old image
		if product.Image != "" {
			oldImagePath := filepath.Join(imageFolder, product.Image)
			if err := os.Remove(oldImagePath); err != nil {
				fmt.Println("Failed to delete old image:", err)
			}
		}

		product.Image = filename
	}

	product.Title = input.Title
	product.Content = input.Content
	product.Price = input.Price
	product.Stock = input.Stock

	if err := pc.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Product updated successfully",
		"data":    product,
	})
}
