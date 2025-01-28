package main

import (
	"fmt"
	"math/big"
	"strconv"
	"sync"
)

// Worker to calculate factorial of numbers from the input channel
func factorialWorker(nums <-chan int, factorials chan<- *big.Int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range nums {
		factorials <- factorial(n)
	}
}

// Worker to calculate the sum of digits of each factorial from the input channel
func digitSumWorker(factorials <-chan *big.Int, digitSums chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for f := range factorials {
		digitSums <- sumOfDigits(f)
	}
}

// Worker to filter digit sums that exceed 50
func filterWorker(digitSums <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for ds := range digitSums {
		if ds > 50 {
			results <- ds
		}
	}
}

// Helper function to calculate factorial using big.Int
func factorial(n int) *big.Int {
	if n < 0 {
		return big.NewInt(-1)
	}
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}

// Helper function to calculate sum of digits of a big.Int
func sumOfDigits(n *big.Int) int {
	digits := n.String()
	sum := 0
	for _, d := range digits {
		num, _ := strconv.Atoi(string(d))
		sum += num
	}
	return sum
}

func main() {
	numbers := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		numbers[i] = i % 20 // Limit factorial complexity
	}

	numWorkers := 10
	nums := make(chan int, numWorkers)
	factorials := make(chan *big.Int, numWorkers)
	digitSums := make(chan int, numWorkers)
	results := make(chan int, numWorkers)

	var wg sync.WaitGroup

	// Stage 1: Factorial computation
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go factorialWorker(nums, factorials, &wg)
	}

	// Stage 2: Sum of digits
	go func() {
		wg.Wait()
		close(factorials)
	}()

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go digitSumWorker(factorials, digitSums, &wg)
	}

	// Stage 3: Filtering digit sums
	go func() {
		wg.Wait()
		close(digitSums)
	}()

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go filterWorker(digitSums, results, &wg)
	}

	// Feeding numbers into the pipeline
	go func() {
		for _, n := range numbers {
			nums <- n
		}
		close(nums)
	}()

	// Final stage: Collect results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Output the results
	for res := range results {
		fmt.Println(res)
	}

	fmt.Println("Pipeline processing complete.")
}