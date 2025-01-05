package main

import (
	"fmt"
	"sync"
	"time"
)

const numWorkers = 100

func processDataParallel(input []int) ([]int, error) {
	wg := &sync.WaitGroup{}
	results := make(chan []int, numWorkers)
	errorsChan := make(chan error, numWorkers)

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(input, results, errorsChan, wg)
	}

	wg.Wait()
	close(results)
	close(errorsChan)

	// Collect and merge results
	var result []int
	for res := range results {
		result = append(result, res...)
	}

	// Check for any errors
	for err := range errorsChan {
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func worker(input []int, results chan<- []int, errors chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	
	// Simulate some work
	time.Sleep(time.Millisecond * 100)
	
	result := make([]int, 0, len(input))
	for _, v := range input {
		result = append(result, v*v)
	}
	
	if len(input) > 100 {
		errors <- errors.New("data size exceeds limit")