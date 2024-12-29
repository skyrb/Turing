package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func copyFile(src, dst string) error {
	// Open the source file.
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close() // Ensure the source file is closed after copying.

	// Create the destination file.
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close() // Ensure the destination file is closed after copying.

	// Log the copying process.
	log.Printf("Copying file from '%s' to '%s'", src, dst)

	// Copy the content from source to destination.
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	src := "source.txt"
	dst := "destination.txt"

	// Measure the execution time of the file copying process.
	start := time.Now()
	defer func() {
		log.Printf("File copying process took: %s", time.Since(start))
	}()

	err := copyFile(src, dst)
	if err != nil {
		log.Fatalf("Error copying file: %v", err)
	}

	fmt.Println("File copied successfully!")
}  