package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

const counterFilePath = "/app/data/pingpong_count.txt"

var (
	counter int64
	mu      sync.Mutex
)

type PingPongResponse struct {
	Pong int64 `json:"pong"`
}

func readCounterFromFile() (int64, error) {
	data, err := os.ReadFile(counterFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, nil // File doesn't exist yet, start at 0
		}
		return 0, err
	}

	var count int64
	_, err = fmt.Sscanf(string(data), "%d", &count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func writeCounterToFile(count int64) error {
	// Ensure directory exists
	dir := filepath.Dir(counterFilePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Write atomically using temp file + rename
	tmpFile := counterFilePath + ".tmp"
	data := fmt.Sprintf("%d", count)

	if err := os.WriteFile(tmpFile, []byte(data), 0644); err != nil {
		return err
	}

	return os.Rename(tmpFile, counterFilePath)
}

func pingPongHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// Increment counter
	counter++
	currentCount := counter

	// Write to persistent file
	if err := writeCounterToFile(currentCount); err != nil {
		fmt.Printf("Error writing counter to file: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

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
	// Load counter from file on startup
	var err error
	counter, err = readCounterFromFile()
	if err != nil {
		fmt.Printf("Error loading counter from file: %v\n", err)
		fmt.Println("Starting with counter = 0")
		counter = 0
	} else {
		fmt.Printf("Loaded counter from file: %d\n", counter)
	}

	// Register the handler
	http.HandleFunc("/pingpong", pingPongHandler)

	fmt.Println("Server starting on :8080")
	fmt.Println("Try: curl http://localhost:8080/pingpong")

	// Start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
