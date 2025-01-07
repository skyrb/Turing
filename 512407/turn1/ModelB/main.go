package main
import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, jobs <-chan int) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, job)
	}
}

func main() {
	numWorkers := 3
	numJobs := 5
	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup
	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, jobs)
	}
	// Close the jobs channel after sending all jobs
	go func() {
		for i := 1; i <= numJobs; i++ {
			jobs <- i
		}
		close(jobs)
	}()
	// Wait for all workers to complete their jobs
	wg.Wait()
	fmt.Println("All jobs done.")
}