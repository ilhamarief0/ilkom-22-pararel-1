package router

import (
	"net/http"
	"order-service/handler"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/orders", handler.CreateOrderHandler).Methods(http.MethodPost)
	r.HandleFunc("/orders", handler.GetOrdersHandler).Methods(http.MethodGet)
	r.HandleFunc("/orders/{id}", handler.GetOrderByIdHandler).Methods(http.MethodGet)
	return r
}
