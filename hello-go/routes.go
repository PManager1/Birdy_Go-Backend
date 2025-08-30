package main

import (
	"encoding/json"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
}

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", PingHandler)
	// Add more routes here
	return mux
}
