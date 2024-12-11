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
	t.Run("Zero", func(t *testing.T) {
		if a := add(0, 0); a != 0 {
			t.Errorf("Expected 0, got %d", a)
		}
	})
	t.Run("LargePositive", func(t *testing.T) {
		if a := add(1<<30, 1); a != (1<<30)+1 {
			t.Errorf("Expected %d, got %d", (1<<30)+1, a)
		}
	})
	t.Run("LargeNegative", func(t *testing.T) {
		if a := add(-1<<30, -1); a != -(1<<30)-1 {
			t.Errorf("Expected %d, got %d", -(1<<30)-1, a)
		}
	})
	t.Run("Overflow", func(t *testing.T) {
		// This test is expected to fail due to integer overflow
		if a := add(1<<31, 1); a != -(1 << 31) {
			t.Errorf("Expected overflow to -%d, got %d", 1<<31, a)
		}
	})
}

func TestMultiply(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		if a := multiply(2, 3); a != 6 {
			t.Errorf("Expected 6, got %d", a)
		}
	})
	t.Run("Negative", func(t *testing.T) {
		if a := multiply(-2, -3); a != 6 {
			t.Errorf("Expected 6, got %d", a)
		}
	})
	t.Run("Zero", func(t *testing.T) {
		if a := multiply(0, 5); a != 0 {
			t.Errorf("Expected 0, got %d", a)
		}
		if a := multiply(5, 0); a != 0 {
			t.Errorf("Expected 0, got %d", a)
		}
	})
	t.Run("LargePositive", func(t *testing.T) {
		if a := multiply(1<<20, 1<<10); a != (1 << 30) {
			t.Errorf("Expected %d, got %d", (1 << 30), a)
		}
	})
	t.Run("LargeNegative", func(t *testing.T) {
		if a := multiply(-1<<20, -1<<10); a != (1 << 30) {
			t.Errorf("Expected %d, got %d", (1 << 30), a)
		}
	})
	t.Run("Overflow", func(t *testing.T) {
		// This test is expected to fail due to integer overflow
		if a := multiply(1<<30, 3); a != 0 {
			t.Errorf("Expected overflow to 0, got %d", a)
		}
	})
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}