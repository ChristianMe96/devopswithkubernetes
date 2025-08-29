package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"
)

var counter int64

type PingPongResponse struct {
	Pong int64 `json:"pong"`
}

func pingPongHandler(w http.ResponseWriter, r *http.Request) {
	// Increment counter atomically (thread-safe)
	currentCount := atomic.AddInt64(&counter, 1)

	// Create response
	response := PingPongResponse{
		Pong: currentCount,
	}

	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode and send JSON response
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Register the handler
	http.HandleFunc("/pingpong", pingPongHandler)

	fmt.Println("Server starting on :8080")
	fmt.Println("Try: curl http://localhost:8080/pingpong")

	// Start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
