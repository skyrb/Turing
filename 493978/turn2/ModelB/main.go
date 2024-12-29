package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func handler(w http.ResponseWriter, r *http.Request) {
	span, ctx := opentracing.StartSpanFromContext(r.Context(), "handler")
	defer span.Finish()

	var wg sync.WaitGroup
	var mu sync.Mutex
	var buffer bytes.Buffer

	// Simulate concurrent processing of a request
	numTasks := 5
	wg.Add(numTasks)

	// Create a context with a timeout for each task
	taskCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	for i := 0; i < numTasks; i++ {
		go func(taskID int) {
			defer wg.Done()
			span, _ := opentracing.StartSpanFromContext(taskCtx, fmt.Sprintf("task %d", taskID))
			defer span.Finish()

			// Simulate work
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

			// Check if the context is canceled due to timeout
			select {
			case <-taskCtx.Done():
				ext.Error.Set(span, true)
				span.LogFields(opentracing.String("event", "timeout"))
				log.Printf("Task %d timed out\n", taskID)
				return
			default:
			}

			// Safely write to buffer
			mu.Lock()
			defer mu.Unlock()
			buffer.WriteString(fmt.Sprintf("Task %d completed\n", taskID))
		}(i)
	}

	// Wait for all tasks to complete or timeout
	wg.Wait()

	// Write response safely after all tasks are complete or timeout
	mu.Lock()
	buffer.WriteString("All tasks completed\n")
	mu.Unlock()
	w.Write(buffer.Bytes())
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	// Initialize tracing (e.g., using Jaeger)
	tracer, closer, err := initTracer("my-web-app")
	if err != nil {
		log.Fatalf("Error initializing tracing: %v\n", err)
	}
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)

	http.HandleFunc("/", handler)
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
 