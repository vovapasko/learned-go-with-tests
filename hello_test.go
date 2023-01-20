package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Wowa")
	want := "Hello, Wowa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
