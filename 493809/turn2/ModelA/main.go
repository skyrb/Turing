package main

import (
	"fmt"
	"log"
	"strings"
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

// Log logs a message with the specified format, handling multi-line messages
func (l *Logger) Log(level string, message interface{}) {
	timestamp := time.Now().Format(time.RFC3339Nano)
	
	// Convert message to a string
	messageString := fmt.Sprint(message)
	
	// Split the message into lines
	lines := strings.Split(messageString, "\n")
	
	// Log each line prefixed with timestamp and log level
	for _, line := range lines {
		formattedLine := fmt.Sprintf(l.Format, timestamp, line)
		log.Print(formattedLine)
	}
}

func main() {
	logger := NewLogger()

	// Log an error with a stack trace
	func test() {
		panic("This is a panic!")
	}
	defer func() {
		if r := recover(); r != nil {
			logger.Log("ERROR", fmt.Sprintf("Panic recovered: %v\n%s", r, recover()))
		}
	}()
	test()

	// Log a warning with a detailed error context
	logger.Log("WARNING", "Invalid input detected: {'name': 'John', 'email': 'invalid@example'}")

	// Log an informational message with multiple lines
	logger.Log("INFO", "Application started successfully:\n  Version: 1.2.3\n  Configuration loaded from /etc/config.json\n  Databases connected: db1, db2")
}