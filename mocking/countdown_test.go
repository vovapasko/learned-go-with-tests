package mocking

import (
	"bytes"
	"testing"
)

func TestCounting(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}
	Counting(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
	if spySleeper.CallsAmount != 3 {
		t.Errorf("Got %d sleeps, want %d sleeps", spySleeper.CallsAmount, 3)
	}

}
