package controllers

import (
	"errors"
	"net/http"
	"product_service/backend-api/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidatePostInput struct {
	Title   string `json:"title" binding :"required"`
	Content string `json:"content" binding : "required"`
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This Field is Required"
	}
	return "Unknown error"
}

func Findpost(c *gin.Context) {
	var posts []models.Product
	models.DB.Find(&posts)

	c.JSON(200, gin.H{
		"succes":   true,
		"messaage": "list data product",
		"data":     posts,
	})
}

func StorePost(c *gin.Context) {
	var input ValidatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}
	post := models.Product{
		Title:   input.Title,
		Content: input.Content,
	}
	models.DB.Create(&post)
	c.JSON(201, gin.H{
		"succes":  true,
		"message": "post Created succesfully",
		"data":    post,
	})
}
