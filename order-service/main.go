package main

import (
	"log"
	"net/http"
	"order-service/router"
)

func main() {
	r := router.SetupRouter()

	log.Println("Order Service is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
