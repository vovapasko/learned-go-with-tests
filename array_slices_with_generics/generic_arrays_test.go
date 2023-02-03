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
		AssertEqual(t, actual, expected)

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

func TestBadBank(t *testing.T) {
	transactions := []Transaction{
		{
			From: "Chris",
			To:   "Riya",
			Sum:  100,
		},
		{
			From: "Adil",
			To:   "Chris",
			Sum:  25,
		},
	}

	AssertEqual(t, BalanceFor(transactions, "Riya"), 100)
	AssertEqual(t, BalanceFor(transactions, "Chris"), -75)
	AssertEqual(t, BalanceFor(transactions, "Adil"), -25)
}

func AssertEqual[T comparable](t *testing.T, want, got T) {
	if want != got {
		t.Errorf("Actual %v, but expected %v", got, want)
	}
}
