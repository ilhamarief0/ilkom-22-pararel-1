package handler

import (
	"encoding/json"
	"net/http"
	"order-service/model"
	"order-service/repository"

	"github.com/google/uuid"
)

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order model.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	order.ID = uuid.New().String()
	repository.CreateOrder(order)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	orders := repository.GetOrders()
	json.NewEncoder(w).Encode(orders)
}

func GetOrderByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	order, err := repository.GetOrderById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(order)
}
