package mathops

import "errors"

// Add adds two integers and returns an error if the result is invalid.
func Add(a, b int) (int, error) {
    if a < 0 || b < 0 {
        return 0, errors.New("negative numbers are not allowed")
    }
    return a + b, nil
}

// Multiply multiplies two integers.
func Multiply(a, b int) int {
    return a * b
}
