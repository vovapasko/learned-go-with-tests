package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("test assertions with integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})
}

func AssertNotEqual(t *testing.T, got int, want int) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %d", want)
	}
}

func AssertEqual(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
