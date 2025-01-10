package handlers

import (
	"encoding/json"
	"net/http"
)

type CreateOrderRequest struct {
	Product  string  `json:"product"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	order := map[string]interface{}{
		"id":       "3", // ID biasanya di-generate otomatis (misalnya menggunakan UUID).
		"product":  req.Product,
		"quantity": req.Quantity,
		"price":    req.Price,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
