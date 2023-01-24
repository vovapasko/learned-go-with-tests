package arrays_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("run with array size 5", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		actual := Sum(numbers)
		expected := 15
		if actual != expected {
			t.Errorf("Actual %d, but expected %d, given %v", actual, expected, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	actual := SumAll([]int{1, 2}, []int{3, 9})
	expected := []int{3, 12}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
