package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"
	"sync"
	"io"
)

func handleConnectionWithContext(ctx context.Context, conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure wg is decremented when done
	defer conn.Close() // Ensure connection is closed

	select {
	case <-ctx.Done():
		log.Println("Connection cancelled by context.")
		return
	default:
		fmt.Fprintf(conn, "Hello from server! The time is %s.\n", time.Now())

		buf := make([]byte, 1024)
		for {
			select {
			case <-ctx.Done():
				log.Println("Connection cancelled by context during read.")
				return
			default:
				n, err := conn.Read(buf)
				if err != nil {
					if err == io.EOF {
						fmt.Println("Client disconnected.")
						break
					}
					log.Printf("Error reading from connection: %v", err)
					return
				}
				fmt.Printf("Received: %s\n", string(buf[:n]))
			}
		}
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

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel() // Ensure context is cancelled if not already cancelled

		wg.Add(1)
		go handleConnectionWithContext(ctx, conn, &wg)
	}
}