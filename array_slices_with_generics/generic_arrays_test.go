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
	var (
		riya  = Account{Name: "Riya", Balance: 100}
		chris = Account{Name: "Chris", Balance: 75}
		adil  = Account{Name: "Adil", Balance: 200}

		transactions = []Transaction{
			NewTransaction(chris, riya, 100),
			NewTransaction(adil, chris, 25),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	AssertEqual(t, newBalanceFor(riya), 200)
	AssertEqual(t, newBalanceFor(chris), 0)
	AssertEqual(t, newBalanceFor(adil), 175)
}

func AssertEqual[T comparable](t *testing.T, want, got T) {
	if want != got {
		t.Errorf("Actual %v, but expected %v", got, want)
	}
}
