package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Tambahkan endpoint sederhana
	r.HandleFunc("/orders", GetOrders).Methods("GET")
	r.HandleFunc("/orders", CreateOrder).Methods("POST")

	fmt.Println("Order Service running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}

// Handler untuk mendapatkan daftar pesanan
func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("List of Orders"))
}

// Handler untuk membuat pesanan baru
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Order Created"))
}
