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
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("failed to close file: %v", err)
		}
	}()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)
	// Defer closure of the scanner's buffer
	defer func() {
		if err := scanner.Buffer().Close(); err != nil {
			log.Fatalf("failed to close scanner buffer: %v", err)
		}
	}()

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