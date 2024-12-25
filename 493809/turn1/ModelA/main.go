package main

import (
	"fmt"
	"log"
	"time"
)

// Logger contains the configuration for log formatting
type Logger struct {
	Format string // Format for the log message
}

// NewLogger initializes a new Logger with the default format
func NewLogger() *Logger {
	return &Logger{Format: "%s: %v\n"}
}

// Log logs a message with the specified format
func (l *Logger) Log(level string, message interface{}) {
	timestamp := time.Now().Format(time.RFC3339Nano)
	formattedMessage := fmt.Sprintf(l.Format, timestamp, message)
	log.Print(formattedMessage)
}
func main() {
	logger := NewLogger()

	// Log an error
	logger.Log("ERROR", "Failed to connect to database: could not find server")

	// Log a warning
	logger.Log("WARNING", "Disk space usage is above 80%")

	// Log an informational message
	logger.Log("INFO", "Application started successfully")

	// Log a debug message
	logger.Log("DEBUG", "Memory usage: 200MB")
}