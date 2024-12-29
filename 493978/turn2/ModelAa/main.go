package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Level     string    `json:"level"`
	Message   string    `json:"message"`
	TaskID    int       `json:"taskID"`
}

func logStructured(level string, message string, taskID int) {
	entry := LogEntry{
		Timestamp: time.Now(),
		Level:     level,
		Message:   message,
		TaskID:    taskID,
	}
	data, err := json.Marshal(entry)
	if err != nil {
		log.Printf("Error marshalling log entry: %v\n", err)
		return
	}
	log.Println(string(data))
}

func handler(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var buffer bytes.Buffer

	// Simulate concurrent processing of a request
	numTasks := 5
	wg.Add(numTasks)

	for i := 0; i < numTasks; i++ {
		go func(taskID int) {
			defer wg.Done()

			// Create a context with a timeout
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()

			select {
			case <-ctx.Done():
				logStructured("WARN", fmt.Sprintf("Task %d timed out", taskID), taskID)
				return
			default:
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
				mu.Lock()
				buffer.WriteString(fmt.Sprintf("Task %d completed\n", taskID))
				mu.Unlock()
			}
		}(i)
	}

	// Wait for all tasks to complete or time out
	wg.Wait()

	// Write response safely after all tasks are complete
	mu.Lock()
	buffer.WriteString("All tasks completed or timed out\n")
	mu.Unlock()
	w.Write(buffer.Bytes())
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	http.HandleFunc("/", handler)
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}