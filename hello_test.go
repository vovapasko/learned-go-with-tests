package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Testing with given name", func(t *testing.T) {
		got := Hello("Wowa")
		want := "Hello, Wowa"

		runAssertStatement(got, want, t)
	})
	t.Run("testing with empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		runAssertStatement(got, want, t)
	})
}

func runAssertStatement(got string, want string, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
