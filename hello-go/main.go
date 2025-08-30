package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Global MongoDB client
var client *mongo.Client

// Structs
type User struct {
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}

type Trip struct {
	From   string `json:"from" bson:"from"`
	To     string `json:"to" bson:"to"`
	Price  string `json:"price" bson:"price"`
	Rider  string `json:"rider" bson:"rider"`
	Driver string `json:"driver" bson:"driver"`
}

// Initialize MongoDB
func initMongo() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system env variables")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Mongo connect error:", err)
	}

	// Ping
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Mongo ping error:", err)
	}

	fmt.Println("✅ Connected to MongoDB!")
}

// Handlers
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	collection := client.Database("test").Collection("users") // Replace 'testdb' with your DB name
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, map[string]interface{}{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var users []User
	for cursor.Next(ctx) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getTripsHandler(w http.ResponseWriter, r *http.Request) {
	collection := client.Database("testdb").Collection("trips") // Replace 'testdb' with your DB name
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, map[string]interface{}{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var trips []Trip
	for cursor.Next(ctx) {
		var trip Trip
		if err := cursor.Decode(&trip); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		trips = append(trips, trip)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trips)
}

func main() {
	initMongo()

	http.HandleFunc("/users", getUsersHandler)
	http.HandleFunc("/trips", getTripsHandler)

	fmt.Println("Server running on http://localhost:3030")
	if err := http.ListenAndServe(":3030", nil); err != nil {
		log.Fatal(err)
	}
}
