package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
	"io"
)

func handleConnection(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure wg is decremented when done
	defer conn.Close() // Ensure connection is closed

	fmt.Fprintf(conn, "Hello from server! The time is %s.\n", time.Now())

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected.")
				break
			}
			log.Printf("Error reading from connection: %v", err)
			break
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

	var wg sync.WaitGroup

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		wg.Add(1)
		go handleConnection(conn, &wg)
	}
}