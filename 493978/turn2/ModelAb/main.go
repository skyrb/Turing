package main

import (
	"context"
	"io"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
)

var tracer opentracing.Tracer

func init() {
	// Initialize the tracer
	// Replace with your actual tracer configuration
	tracer, err := opentracing.NewTracer("my-tracer")
	if err != nil {
		log.Fatalf("Failed to initialize tracer: %v\n", err)
	}
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

			// Create a context with a span
			ctx, span := opentracing.StartSpanFromContext(context.Background(), "task")
			defer span.Finish()

			span.LogFields(
				ext.Error(false),
				ext.Message("Starting task"),
				log.Kv("taskID", taskID),
			)

			// Create a context with a timeout
			ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
			defer cancel()

			select {
			case <-ctx.Done():
				logStructured("WARN", fmt.Sprintf("Task %d timed out", taskID), taskID)
				span.LogFields(
					ext.Error(true),
					ext.Message("Task timed out"),
					log.Kv("taskID", taskID),
				)
				return
			default:
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
				mu.Lock()
				buffer.WriteString(fmt.Sprintf("Task %d completed\n", taskID))
				mu.Unlock()
				span.LogFields(
					ext.Error(false),
					ext.Message("Task completed"),
					log.Kv("taskID", taskID),
				)
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