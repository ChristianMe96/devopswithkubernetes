package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

const logFilePath = "/app/logs/random.log"
const pingPongCounterPath = "/app/data/pingpong_count.txt"
const port = ":8080"

// LogResponse represents the JSON response structure
type LogResponse struct {
	Logs      []string `json:"logs"`
	Count     int      `json:"count"`
	PingPongs int64    `json:"pingpongs"`
}

// ErrorResponse represents an error in JSON format
type ErrorResponse struct {
	Error string `json:"error"`
}

// readLogFile reads the entire log file and returns its contents as a slice of strings
func readLogFile() ([]string, error) {
	file, err := os.Open(logFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("log file not found - waiting for logs to be generated")
		}
		return nil, fmt.Errorf("failed to open log file: %v", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading log file: %v", err)
	}

	return lines, nil
}

// readPingPongCounter reads the current pingpong counter from the shared volume
func readPingPongCounter() (int64, error) {
	data, err := os.ReadFile(pingPongCounterPath)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, nil // File doesn't exist yet, return 0
		}
		return 0, fmt.Errorf("failed to read counter file: %v", err)
	}

	var count int64
	_, err = fmt.Sscanf(string(data), "%d", &count)
	if err != nil {
		return 0, fmt.Errorf("failed to parse counter: %v", err)
	}

	return count, nil
}

func logOutputHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Read all logs (or use readLastNLines(10) for last 10 lines)
	lines, err := readLogFile()
	if err != nil {
		// Return error message with 503 status if file not ready
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	// Read pingpong counter
	pingPongCount, err := readPingPongCounter()
	if err != nil {
		fmt.Printf("Warning: Could not read pingpong counter: %v\n", err)
		// Continue with count = 0 rather than failing the request
		pingPongCount = 0
	}

	// Create response with separate pingpong count
	response := LogResponse{
		Logs:      lines,
		Count:     len(lines),
		PingPongs: pingPongCount,
	}

	// Encode and send JSON response
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Ensure log directory exists
	logDir := filepath.Dir(logFilePath)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		fmt.Printf("Warning: Could not create log directory: %v\n", err)
	}

	// Register handlers
	http.HandleFunc("/", logOutputHandler)

	fmt.Printf("ReadLog server starting on %s\n", port)
	fmt.Printf("Reading logs from: %s\n", logFilePath)

	// Start the server
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
