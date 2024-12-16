package models

import "gorm.io/gorm"

// Payment represents a payment record
type Payment struct {
	gorm.Model
	OrderID   string // Add this field to store the order ID
	UserID    uint
	ProductID uint
	Amount    float64
	Status    string
}
