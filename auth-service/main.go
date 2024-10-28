package main

import (
	"auth-service/controllers"
	"auth-service/db"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	db.Init() // Initialize database connection

	// Inisialisasi router
	r := mux.NewRouter()

	// Definisikan route
	r.HandleFunc("/api/auth/login", controllers.Login).Methods("POST")
	r.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/user", controllers.CreateUser).Methods("POST")

	// Set up CORS handler
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4000"}, // Izinkan origin ini
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// Jalankan server dengan middleware CORS
	handler := corsHandler.Handler(r)
	log.Fatal(http.ListenAndServe(":3012", handler))
}
