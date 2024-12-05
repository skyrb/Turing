package main

import (
	"fmt"
	"log"
	"os"
)

func handleMultipleFiles(filePaths []string) {
	var files []*os.File

	for _, path := range filePaths {
		// Open the file and ensure it's added to the list of files to be closed
		file, err := os.Open(path)
		if err != nil {
			log.Printf("failed to open file %s: %v", path, err)
			continue // Skip this file and proceed to next
		}

		// Defer the closing of each file, capturing any errors
		defer func(f *os.File) {
			if err := f.Close(); err != nil {
				log.Printf("failed to close file %s: %v", f.Name(), err)
			}
		}(file)

		files = append(files, file)
	}

	// Process files here (assuming we have non-empty slice of files)
	for _, file := range files {
		fmt.Printf("Processing file: %s\n", file.Name())
		// Add logic to process each file
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	filePaths := []string{"file1.txt", "file2.txt", "file3.txt"}
	handleMultipleFiles(filePaths)
}