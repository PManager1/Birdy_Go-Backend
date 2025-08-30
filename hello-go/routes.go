package main

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

// PingHandler responds with "pong"
func PingHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
}

// HomeHandler responds to "/" requests
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is working!"))
}

// UsersHandler fetches all users from MongoDB
func UsersHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		collection := db.Collection("users") // collection name from Node: users

		cursor, err := collection.Find(context.Background(), struct{}{})
		if err != nil {
			http.Error(w, "Error fetching users", http.StatusInternalServerError)
			return
		}
		defer cursor.Close(context.Background())

		var users []map[string]interface{}
		if err := cursor.All(context.Background(), &users); err != nil {
			http.Error(w, "Error decoding users", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(users)
	}
}

// RegisterRoutes sets up routes and injects MongoDB database
func RegisterRoutes(db *mongo.Database) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is working!"))
	})
	mux.HandleFunc("/ping", PingHandler)
	mux.HandleFunc("/users", UsersHandler(db))
	return mux
}
