package example

import (
	"testing"
)

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
}

func TestMultiply(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		if a := multiply(2, 3); a != 6 {
			t.Errorf("Expected 6, got %d", a)
		}
	})
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}