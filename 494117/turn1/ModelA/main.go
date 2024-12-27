package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	// Example of using defer to release a mutex lock
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()

	fmt.Println("Mutex locked.")

	// Example of using defer to close a file
	file, err := os.Create("temp.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Println("File created.")

	// Write some data to the file
	_, err = file.WriteString("Hello, World!")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Data written to file.")
}