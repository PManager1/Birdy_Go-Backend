package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	godotenv.Load()

	// Call RegisterRoutes directly, no import needed
	router := RegisterRoutes()

	port := "3030"
	log.Println("Server running on port", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
