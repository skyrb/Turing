package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

const maxConnections = 100
const batchSize = 10

func handleConnection(ctx context.Context, conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure wg is decremented when done

	fmt.Fprintf(conn, "Hello from server! The time is %s.\n", time.Now())

	buf := make([]byte, 1024)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Client disconnected (context canceled).")
			return
		default:
		}

		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected.")
				return
			}
			log.Printf("Error reading from connection: %v", err)
			return
		}
		fmt.Printf("Received: %s\n", string(buf[:n]))
	}
}

func closeConnections(connections []net.Conn) {
	for _, conn := range connections {
		conn.Close()
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close() // Ensure listener is closed

	fmt.Println("Server is listening on :8080")

	var wg sync.WaitGroup
	var connections []net.Conn
	var openCount atomic.Int32

	// Create a context that can be canceled later
	ctx, cancel := context.WithCancel(context.Background())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		openCount.Inc()
		connections = append(connections, conn)

		if openCount.Load() % batchSize == 0 {
			go closeConnections(connections[:batchSize])
			connections = connections[batchSize:]
		}

		wg.Add(1)
		go handleConnection(ctx, conn, &wg)
	}

	// To simulate a dynamic connection lifecycle,
	// we can cancel the context after a certain duration.
	time.AfterFunc(5*time.Second, func() {
		cancel()
		fmt.Println("Context canceled, shutting down connections...")
	})

	// Wait for all goroutines to finish
	wg.Wait()
	closeConnections(connections)
	fmt.Println("All connections closed.")
}