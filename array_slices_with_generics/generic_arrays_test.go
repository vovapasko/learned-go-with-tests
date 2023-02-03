package arrays_slices_with_generics

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("run reduce addition on array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		addFunction := func(a, b int) int { return a + b }
		actual := Reduce(numbers, addFunction, 0)
		expected := 15
		if actual != expected {
			t.Errorf("Actual %d, but expected %d, given %v", actual, expected, numbers)
		}
	})
	t.Run("run reduce multiplication on array", func(t *testing.T) {
		numbers := []float32{1, 2, 3, 4, 5}
		multiplyFunction := func(a, b float32) float32 { return a * b }
		actual := Reduce(numbers, multiplyFunction, 1)
		expected := float32(120)
		if actual != expected {
			t.Errorf("Actual %v, but expected %v, given %v", actual, expected, numbers)
		}
	})
}
