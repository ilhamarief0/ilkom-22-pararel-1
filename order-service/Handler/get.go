package handlers

import (
	"encoding/json"
	"net/http"
)

type Order struct {
	ID       string  `json:"id"`
	Product  string  `json:"product"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	orders := []Order{
		{ID: "1", Product: "Laptop", Quantity: 2, Price: 2000.00},
		{ID: "2", Product: "Mouse", Quantity: 5, Price: 50.00},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
