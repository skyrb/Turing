package main

import (
	"container/heap"
	"fmt"
	"math/big"
	"strconv"
	"sync"
)

// Task represents a number with a priority
type Task struct {
	Number   int
	Priority int
}

// PriorityQueue implements a priority queue using a heap
type PriorityQueue []*Task

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// Higher priority comes first
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Task)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Worker to calculate factorial of numbers from the task channel
func factorialWorker(tasks <-chan *Task, factorials chan<- *big.Int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		factorials <- factorial(task.Number)
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
	// Create 100,000 tasks with pseudo-random priorities
	tasks := make(PriorityQueue, 100000)
	for i := range tasks {
		tasks[i] = &Task{
			Number:   i % 20,              // Limiting for factorial calculation
			Priority: 100000 - (i % 100), // Assigning priorities for demonstration
		}
	}
	heap.Init(&tasks)

	numWorkers := 10
	taskChan := make(chan *Task, numWorkers)
	factorials := make(chan *big.Int, numWorkers)
	digitSums := make(chan int, numWorkers)
	results := make(chan int, numWorkers)

	var factorialWg sync.WaitGroup
	var digitSumWg sync.WaitGroup
	var filterWg sync.WaitGroup

	// Stage 1: Factorial computation
	for i := 0; i < numWorkers; i++ {
		factorialWg.Add(1)
		go factorialWorker(taskChan, factorials, &factorialWg)
	}

	// Stage 2: Sum of digits
	go func() {
		factorialWg.Wait()
		close(factorials)
	}()

	for i := 0; i < numWorkers; i++ {
		digitSumWg.Add(1)
		go digitSumWorker(factorials, digitSums, &digitSumWg)
	}

	// Stage 3: Filtering digit sums
	go func() {
		digitSumWg.Wait()
		close(digitSums)
	}()

	for i := 0; i < numWorkers; i++ {
		filterWg.Add(1)
		go filterWorker(digitSums, results, &filterWg)
	}

	// Feeding tasks into the pipeline based on priority
	go func() {
		for tasks.Len() > 0 {
			task := heap.Pop(&tasks).(*Task)
			taskChan <- task
		}
		close(taskChan)
	}()

	// Final stage: Collect results
	go func() {
		filterWg.Wait()
		close(results)
	}()

	// Output the results
	for res := range results {
		fmt.Println(res)
	}

	fmt.Println("Pipeline processing complete.")
}