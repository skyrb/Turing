package main  
import (  
    "fmt"
    "sync"
)
const workers = 100 // Number of goroutines for parallel processing

// FactorialFunction calculates factorial of the given number
func FactorialFunction(in <-chan int, out chan<- int) {  
    for num := range in {  
        if num < 0 {
            out <- 0
        } else if num == 0 || num == 1 {
            out <- 1
        } else {
            result := 1
            for i := 1; i <= num; i++ {
                result *= i
            }
            out <- result
        }
    }  
}  
func main() {  
    nums := make([]int, 1_000_000)  
    for i := range nums {
        nums[i] = i
    }
    results := make([]int, len(nums))  
    // Create buffered channels for data interchange between workers and the main goroutine
    in := make(chan int, len(nums))  
    out := make(chan int, len(nums))  
    // Spawn a pool of goroutines to perform factorial calculations
    var wg sync.WaitGroup  
    wg.Add(workers)  
    for i := 0; i < workers; i++ {  
        go func() {  
            FactorialFunction(in, out)  
            wg.Done()  
        }()  
    }  
    // Pass the input data to the workers via the in channel
    go func() {  
        for _, num := range nums {  
            in <- num  
        }  
        close(in)  
    }()  
    // Collect the results from the workers via the out channel
    go func() {  
        for i := range out {  
            results[i] = i
        }  
    }()  
    // Wait for all workers to complete their tasks
    wg.Wait()  
    // Close the out channel to signal the end of data
    close(out)  
    // Print the results
    for _, result := range results {  
        fmt.Println(result)
    }  
}  
