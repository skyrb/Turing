package main  
import (  
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

// calculateSum logs and calculates the sum of the integers in a slice.
func calculateSum(numbers []int) int {
	log.Printf("Calculating sum: %v", numbers)
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// calculateAverage logs and calculates the average of the integers in a slice.
func calculateAverage(numbers []int) float64 {
	log.Printf("Calculating average: %v", numbers)
	sum := calculateSum(numbers)
	if len(numbers) == 0 {
		return 0
	}
	return float64(sum) / float64(len(numbers))
}

// filterEven logs and filters even numbers from a slice of integers.
func filterEven(numbers []int) []int {
	var evens []int
	for _, num := range numbers {
		if num%2 == 0 {
			evens = append(evens, num)
		}
	}
	log.Printf("Filtered even numbers: %v", evens)
	return evens
}

// reverseSlice logs and reverses a slice of integers.
func reverseSlice(numbers []int) []int {
	reversed := make([]int, len(numbers))
	copy(reversed, numbers)
	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}
	log.Printf("Reversed slice: %v", reversed)
	return reversed
}

// parseInput converts a slice of strings to a slice of integers.
func parseInput(args []string) ([]int, error) {
	var numbers []int
	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			return nil, fmt.Errorf("invalid number: %s", arg)
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Println("What operation do you want to perform?")
        fmt.Println("1. Calculate sum")
        fmt.Println("2. Calculate average")
        fmt.Println("3. Filter even numbers")
        fmt.Println("4. Reverse slice")
        fmt.Println("5. Quit")
        fmt.Print("Enter your choice: ")
        scanner.Scan()
        choice := scanner.Text()

        switch choice {
        case "1", "sum":
            fmt.Println("Please enter the list of numbers:")
            scanner.Scan()
            input := scanner.Text()
            processSum(input)
        case "2", "average":
            fmt.Println("Please enter the list of numbers:")
            scanner.Scan()
            input := scanner.Text()
            processAverage(input)          
        case "3", "filterEven":
            fmt.Println("Please enter the list of numbers:")
            scanner.Scan()
            input := scanner.Text()
            processFilterEven(input)          
        case "4", "reverse":
            fmt.Println("Please enter the list of numbers:")
            scanner.Scan()
            input := scanner.Text()
            processReverse(input)
        case "5", "quit":
            fmt.Println("Goodbye!")
            return
        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}

func processSum(input string) {
    numbers, err := parseInput(strings.Fields(input))
    if err != nil {
        log.Fatalf("Error parsing numbers: %v", err)
    }
    sum := calculateSum(numbers)
    fmt.Printf("Sum: %d\n", sum)
}

func processAverage(input string) {    
    numbers, err := parseInput(strings.Fields(input))
    if err != nil {
        log.Fatalf("Error parsing numbers: %v", err)
    }
    average := calculateAverage(numbers)
    fmt.Printf("Average: %.2f\n", average)
}

func processFilterEven(input string) {    
    numbers, err := parseInput(strings.Fields(input))
    if err != nil {
        log.Fatalf("Error parsing numbers: %v", err)
    }
    evens := filterEven(numbers)
    fmt.Printf("Filtered Even Numbers: %v\n", evens)
}

func processReverse(input string) {   
    numbers, err := parseInput(strings.Fields(input))
    if err != nil {
        log.Fatalf("Error parsing numbers: %v", err)
    }
    reversed := reverseSlice(numbers)
    fmt.Printf("Reversed Slice: %v\n", reversed)
} 

// Rest of the functions (calculateSum, calculateAverage, filterEven, reverseSlice, parseInput) remain the same