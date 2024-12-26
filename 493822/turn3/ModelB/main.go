package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

const (
	maxConnections = 100
	connectionCloseTimeout = 5 * time.Second
)

var (
	connectionPool sync.Pool
	connectionWaitGroup sync.WaitGroup
)

type connection struct {
	net.Conn
}

func init() {
	connectionPool.New = func() interface{} {
		return &connection{}
	}
}

func handleConnection(ctx context.Context, conn net.Conn) {
	defer connectionWaitGroup.Done()

	// Use the connection from the pool
	c := connectionPool.Get().(*connection)
	c.Conn = conn
	defer func() {
		// Return the connection to the pool
		connectionPool.Put(c)
	}()

	fmt.Fprintf(c, "Hello from server! The time is %s.\n", time.Now())

	buf := make([]byte, 1024)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Client disconnected (context canceled).")
			return
		default:
		}

		n, err := c.Read(buf)
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

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close() // Ensure listener is closed

	fmt.Println("Server is listening on :8080")

	// Create a context that can be canceled later
	ctx, cancel := context.WithCancel(context.Background())