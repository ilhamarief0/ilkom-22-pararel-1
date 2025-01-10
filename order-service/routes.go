package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilhamarief0/ilkom-22-pararel-1/services/order/handlers"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/orders", handlers.GetOrders).Methods(http.MethodGet)
	r.HandleFunc("/orders", handlers.CreateOrder).Methods(http.MethodPost)
}
