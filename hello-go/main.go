// package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/joho/godotenv"
// )

// func main() {
// 	// Load .env
// 	godotenv.Load()

// 	// Call RegisterRoutes directly, no import needed
// 	// router := RegisterRoutes()
// 	// router := routes.RegisterRoutes(mongoClient.Database("testdb"))
// 	router := RegisterRoutes(mongoClient.Database("test"))

// 	port := "3030"
// 	log.Println("Server running on port", port)
// 	if err := http.ListenAndServe(":"+port, router); err != nil {
// 		log.Fatal(err)
// 	}
// }

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	godotenv.Load()

	// Create MongoDB client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI not set")
	}

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	// Ping the database to confirm connection
	if err := mongoClient.Ping(ctx, nil); err != nil {
		log.Fatal("MongoDB ping failed:", err)
	}
	log.Println("✅ MongoDB connected")

	// Pass the database to your routes
	router := RegisterRoutes(mongoClient.Database("test"))

	port := "3030"
	log.Println("Server running on port", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
