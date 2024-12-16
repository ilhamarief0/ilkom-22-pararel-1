package handlers

import (
	"fmt"
	"net/http"
	"payment-service/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
)

type PaymentHandler struct {
	DB *gorm.DB
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var payment models.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	timestamp := time.Now().UnixNano()
	// Generate a shorter unique order ID using timestamp
	orderID := fmt.Sprintf("%d", timestamp)

	// Initialize Midtrans Snap
	s := snap.Client{}
	s.New("SB-Mid-server-Dh8ojLkmziWFkCqiJGiDAkq0", midtrans.Sandbox) // Replace with your actual server key

	// Create Snap transaction
	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: int64(payment.Amount),
		},
	}

	snapResp, err := s.CreateTransaction(snapReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	// Check transaction status
	if snapResp != nil && snapResp.StatusCode == "201" {
		payment.Status = "success"
	} else {
		payment.Status = "failed"
	}

	// Save payment to DB
	payment.OrderID = orderID
	h.DB.Create(&payment)

	if payment.Status == "success" {
		c.JSON(http.StatusOK, gin.H{"message": "Payment successful", "snap_redirect_url": snapResp.RedirectURL})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Payment failed"})
	}
}
