package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile(filePath string) error {
	// Attempt to open the file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	// Ensure the file is closed when the function exits
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("failed to close file: %v", err)
		}
	}()

	// Read the file content
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	return nil
}

func main() {
	// Configure logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err := readFile("example.txt"); err != nil {
		log.Printf("Error: %v", err)
	}
}