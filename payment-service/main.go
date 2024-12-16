package main

import (
	"payment-service/db"
	"payment-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	database := db.ConnectDB()

	// Run migrations
	db.Migrate(database)

	// Set up Gin
	r := gin.Default()

	// Initialize handlers
	paymentHandler := handlers.PaymentHandler{DB: database}
	webhookHandler := handlers.WebhookHandler{DB: database}

	// Register routes
	r.POST("/payments", paymentHandler.CreatePayment)
	r.POST("/webhook", webhookHandler.HandleWebhook) // Webhook endpoint

	// Start the server
	r.Run(":8080")
}
