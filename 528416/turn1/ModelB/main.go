
package main  
import (  
    "log"
)

func sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func average(numbers []int) float64 {
	total := float64(sum(numbers))
	return total / float64(len(numbers))
}

func main() {
    log.SetPrefix("Calculating: ")  
    numbers := []int{5, 10, 15, 20}
    
    log.Printf("Sum of %v: %d\n", numbers, sum(numbers))  
    log.Printf("Average of %v: %0.2f\n", numbers, average(numbers))  

}
