package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)

	// Use crypto/rand for truly random bytes
	_, err := rand.Read(b)
	if err != nil {
		// Fallback to timestamp-based string if crypto/rand fails
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}

	// Map random bytes to charset
	for i := range b {
		b[i] = charset[b[i]%byte(len(charset))]
	}

	return string(b)
}

func writeToLogFile(message string) error {
	logDir := "/app/logs"
	err := os.MkdirAll(logDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create log directory: %v", err)
	}

	logFile := filepath.Join(logDir, "random.log")
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}

	defer file.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	_, err = file.WriteString(fmt.Sprintf("[%s] %s\n", timestamp, message))
	if err != nil {
		return fmt.Errorf("failed to write to log file: %v", err)
	}

	return nil
}

func main() {
	// Generate random string once on startup
	randomString := generateRandomString(16)

	fmt.Printf("Starting log generator with random string: %s\n", randomString)
	fmt.Println("Writing to /app/logs/random.log every 5 seconds")

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := writeToLogFile(randomString)
			if err != nil {
				fmt.Printf("Error writing to log file: %v\n", err)
			} else {
				fmt.Printf("Logged: %s\n", randomString)
			}
		}
	}
}
