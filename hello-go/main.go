package main

import (
	"encoding/json"
	"net/http"
)

// Struct for JSON responses
type Response struct {
	Message string `json:"message"`
}

func main() {
	// Register endpoints
	http.HandleFunc("/ping", pingHandler)   // <-- call the function here
	http.HandleFunc("/hello", helloHandler) // <-- call another function

	// Start the server
	http.ListenAndServe(":8080", nil)
}

// Function definitions come after main (or before, both work)

// Handles /ping
func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Message: "pong"})
}

// Handles /hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Message: "Hello, Go!"})
}
