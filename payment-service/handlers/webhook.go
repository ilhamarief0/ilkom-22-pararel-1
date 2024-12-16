package handlers

import (
	"log"
	"net/http"
	"payment-service/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type WebhookHandler struct {
	DB *gorm.DB
}

func (h *WebhookHandler) HandleWebhook(c *gin.Context) {
	var notification map[string]interface{}
	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	orderID, ok := notification["order_id"].(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	transactionStatus, ok := notification["transaction_status"].(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction status"})
		return
	}

	// Update payment status in the database
	var payment models.Payment
	if err := h.DB.Where("order_id = ?", orderID).First(&payment).Error; err != nil {
		log.Printf("Payment not found: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	if transactionStatus == "capture" || transactionStatus == "settlement" {
		payment.Status = "success"
	} else if transactionStatus == "deny" || transactionStatus == "expire" || transactionStatus == "cancel" {
		payment.Status = "failed"
	}

	h.DB.Save(&payment)
	c.JSON(http.StatusOK, gin.H{"message": "Payment status updated"})
}
