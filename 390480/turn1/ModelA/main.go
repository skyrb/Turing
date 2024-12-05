package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func openFile() (error) {
	file, err := os.Open("example.txt")
	if err != nil {
		return err
	}

	defer file.Close() // Ensure the file is closed regardless of the outcome
	
	// Processing the file here
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	fmt.Println(string(contents))
	return nil
}

func main() {
	if err := openFile(); err != nil {
		fmt.Println("Error opening file:", err)
	}
}