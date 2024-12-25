package main  
import (  
    "fmt"
    "log"
    "runtime"
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

// Log logs a message with the specified format.
// If the message is a string containing newline characters (\n),
// it will be logged as a multi-line log.
func (l *Logger) Log(level string, message interface{}) {
    timestamp := time.Now().Format(time.RFC3339Nano)
    formattedMessage := fmt.Sprintf(l.Format, timestamp, message)

    // Check if the message contains newline characters
    if strings.Contains(formattedMessage, "\n") {
        // Add a separator line for multi-line logs
        formattedMessage = "---\n" + formattedMessage + "---\n"
    }

    log.Print(formattedMessage)
}

// LogStackTrace logs a stack trace at the specified log level.
func (l *Logger) LogStackTrace(level string, message string) {
    buf := make([]byte, 1024*10) // 10KB buffer
    runtime.Stack(buf, false)
    l.Log(level, fmt.Sprintf("%s\n%s", message, string(buf)))
}

func main() {
    logger := NewLogger()

    // Log a single-line message
    logger.Log("INFO", "Application started successfully")

    // Log a multi-line message (error context)
    errContext := fmt.Sprintf("Error occurred during processing:\n" +
        "  User ID: 123\n" +
        "  Request Data: %s", "{\"key\": \"value\"}")
    logger.Log("ERROR", errContext)

    // Log a stack trace for an error
    logger.LogStackTrace("ERROR", "Fatal error occurred: panic!")
}