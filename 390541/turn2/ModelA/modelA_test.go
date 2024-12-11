package example

import (
	"testing"
)

// Function definitions
func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

// Test suite
func TestAdd(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		if a := add(1, 2); a != 3 {
			t.Errorf("Expected 3, got %d", a)
		}
	})

	t.Run("Negative", func(t *testing.T) {
		if a := add(-1, -2); a != -3 {
			t.Errorf("Expected -3, got %d", a)
		}
	})

	// Adding edge cases
	t.Run("Zero", func(t *testing.T) {
		if a := add(0, 0); a != 0 {
			t.Errorf("Expected 0, got %d", a)
		}
		if a := add(0, 5); a != 5 {
			t.Errorf("Expected 5, got %d", a)
		}
		if a := add(5, 0); a != 5 {
			t.Errorf("Expected 5, got %d", a)
		}
	})

	t.Run("Large Numbers", func(t *testing.T) {
		if a := add(1000000000, 2000000000); a != 3000000000 {
			t.Errorf("Expected 3000000000, got %d", a)
		}
	})
}

func TestMultiply(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		if a := multiply(2, 3); a != 6 {
			t.Errorf("Expected 6, got %d", a)
		}
	})

	// Adding edge cases
	t.Run("Zero", func(t *testing.T) {
		if a := multiply(0, 100); a != 0 {
			t.Errorf("Expected 0, got %d", a)
		}
		if a := multiply(100, 0); a != 0 {
			t.Errorf("Expected 0, got %d", a)
		}
	})

	t.Run("Negative", func(t *testing.T) {
		if a := multiply(-1, 1); a != -1 {
			t.Errorf("Expected -1, got %d", a)
		}
		if a := multiply(-2, -3); a != 6 {
			t.Errorf("Expected 6, got %d", a)
		}
	})

	t.Run("Large Numbers", func(t *testing.T) {
		if a := multiply(100000, 30000); a != 3000000000 {
			t.Errorf("Expected 3000000000, got %d", a)
		}
	})

	t.Run("Intentional Fail", func(t *testing.T) {
		// This test is intentionally incorrect to demonstrate test failure handling.
		if a := multiply(2, 2); a != 5 {
			t.Errorf("Expected 5, got %d", a) // This will fail
		}
	})
}