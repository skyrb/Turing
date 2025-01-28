package main

import (
	"fmt"
	"math/big"
	"sync"
)

// worker calculates the factorial of numbers received from the nums channel
// and sends the result on the results channel.
func worker(nums <-chan int, results chan<- *big.Int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range nums {
		results <- factorial(n)
	}
}

// factorial calculates the factorial of a number using big.Int.
func factorial(n int) *big.Int {
	if n < 0 {
		return big.NewInt(-1) // Return -1 for negative numbers, which typically don't have factorials
	}
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}

func main() {
	numbers := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		numbers[i] = i % 20 // Modulus to limit factorial complexity
	}

	numWorkers := 10
	nums := make(chan int, numWorkers)         // Channel for input numbers
	results := make(chan *big.Int, numWorkers) // Channel for factorial results

	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(nums, results, &wg)
	}

	// Send numbers to nums channel
	go func() {
		for _, n := range numbers {
			nums <- n
		}
		close(nums) // Close channel when done sending
	}()

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(results) // Close results channel after processing
	}()

	// Collect results (in a real case you might do something more useful with them)
	for result := range results {
		fmt.Println(result)
	}

	fmt.Println("All factorials computed.")
}
