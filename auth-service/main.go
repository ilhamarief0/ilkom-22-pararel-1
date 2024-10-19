package main

import (
	"auth-service/controllers"
	"auth-service/db"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.Init() // Initialize database connection

	r := mux.NewRouter()
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}
