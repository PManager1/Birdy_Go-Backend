// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello, Go!")
// }

package main

import (
	"encoding/json"
	"net/http"
)

type Ping struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Ping{Message: "pong"})
	})
	http.ListenAndServe(":8080", nil)
}
