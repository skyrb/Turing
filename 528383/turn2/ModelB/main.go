package main  
import (  
    "fmt"
    "sync"
)

const (
    numWorkers    = 10
    factorialStage = iota
    digitSumStage
    filterStage
)

// worker performs the specified stage of processing on numbers received from the input channel
// and sends the result on the output channel.
func worker(stage int, input <-chan int, output chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    switch stage {
    case factorialStage:
        for n := range input {
            output <- sumDigits(factorial(n))
        }
    case digitSumStage:
        for num := range input {
            output <- num
        }
    case filterStage:
        for num := range input {
            if num <= 50 {
                output <- num
            }
        }
    }
}

// factorial calculates the factorial of a number using big.Int.
func factorial(n int) int {
    if n < 0 {
        return -1 // Return -1 for negative numbers
    }
    if n == 0 || n == 1 {
        return 1
    }
    result := 1
    for i := 2; i <= n; i++ {
        result *= i
    }
    return result
}

// sumDigits sums the digits of an integer.
func sumDigits(num int) int {
    sum := 0
    for num != 0 {
        digit := num % 10
        sum += digit
        num /= 10
    }
    return sum
}

func main() {
    numbers := make([]int, 1000000)
    for i := 0; i < 1000000; i++ {
        numbers[i] = i % 20
    }

    numStages := 3
    inputChannels := make([]chan int, numStages)
    outputChannels := make([]chan int, numStages)

    var wg sync.WaitGroup

    // Start worker goroutines for each stage
    for stage := 0; stage < numStages; stage++ {
        inputChannels[stage] = make(chan int, numWorkers)
        outputChannels[stage] = make(chan int, numWorkers)
        wg.Add(1)
        go worker(stage, inputChannels[stage], outputChannels[stage], &wg)
    }

    // Start the pipeline by sending numbers to the first stage
    go func() {
        for _, n := range numbers {
            inputChannels[0] <- n
        }
        close(inputChannels[0])
    }()

    // Connect stages
    for stage := 1; stage < numStages; stage++ {
        go func(prevStage, nextStage int) {
            for result := range outputChannels[prevStage] {
                inputChannels[nextStage] <- result
            }
            close(inputChannels[nextStage])
        }(stage-1, stage)
    }

    // Wait for all workers to finish
    go func() {
        wg.Wait()
        close(outputChannels[numStages-1])
    }()

    // Collect results from the last stage (in a real case you might do something more useful with them)
    finalResults := []int{}
    for result := range outputChannels[numStages-1] {
        finalResults = append(finalResults, result)
    }

    fmt.Println("Final results:", finalResults)
    fmt.Println("Processed", len(finalResults), "results")
}  
