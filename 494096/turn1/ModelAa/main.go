package main

import (
	"fmt"
	"os"
)

func processFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close() // Ensure file is closed, even if an error occurs

	// File processing logic goes here
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	fmt.Println(string(data))

	return nil
}

func main() {
	if err := processFile("example.txt"); err != nil {
		fmt.Println("Error:", err)
	}
}