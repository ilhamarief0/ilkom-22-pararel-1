package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"product_service/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidatePostInput struct {
	Title   string `form:"title" json:"title" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
}
type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	}
	return "Unknown error"
}

func FindProduct(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)

	if len(products) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "No products found",
			"data":    products,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "List of products",
		"data":    products,
	})
}

func AddProduct(c *gin.Context) {
	// Mengambil data form
	var input ValidatePostInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	// Mengambil file gambar dari form-data
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
		return
	}

	// Tentukan folder untuk menyimpan gambar
	imageFolder := "gambarproduk"
	if _, err := os.Stat(imageFolder); os.IsNotExist(err) {
		// Membuat folder jika belum ada
		os.Mkdir(imageFolder, os.ModePerm)
	}

	// Tentukan lokasi penyimpanan gambar
	filename := filepath.Base(file.Filename)
	imagePath := filepath.Join(imageFolder, filename)

	// Simpan file gambar ke folder
	if err := c.SaveUploadedFile(file, imagePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	// Membuat produk baru
	product := models.Product{
		Title:   input.Title,
		Content: input.Content,
		Image:   filename, // Simpan nama file gambar di database
	}

	// Simpan data produk ke database
	if err := models.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	// Mengembalikan respons sukses
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Product created successfully",
		"data":    product,
	})
}

func EditProduct(c *gin.Context) {
	// Mendapatkan product berdasarkan id
	var product models.Product
	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found!"})
		return
	}

	// Mengambil data form
	var input ValidatePostInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	// Cek apakah ada file gambar baru yang diunggah
	file, err := c.FormFile("image")
	if err == nil {
		// Tentukan lokasi penyimpanan gambar baru
		imageFolder := "../gambarproduk"
		filename := filepath.Base(file.Filename)
		imagePath := filepath.Join(imageFolder, filename)

		// Simpan gambar baru ke folder
		if err := c.SaveUploadedFile(file, imagePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
			return
		}

		// Hapus gambar lama jika ada
		if product.Image != "" {
			oldImagePath := filepath.Join(imageFolder, product.Image)
			if err := os.Remove(oldImagePath); err != nil {
				fmt.Println("Failed to delete old image:", err)
			}
		}

		// Update nama file gambar di produk
		product.Image = filename
	}

	// Update field lain
	product.Title = input.Title
	product.Content = input.Content

	// Simpan perubahan ke database
	if err := models.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	// Mengembalikan respons sukses
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Product updated successfully",
		"data":    product,
	})
}
