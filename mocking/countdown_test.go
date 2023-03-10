package mocking

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCounting(t *testing.T) {
	t.Run("main test", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &CountdownOperationsSpy{}
		Counting(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}

	})

	t.Run("test operations order", func(t *testing.T) {

		operationsSpy := &CountdownOperationsSpy{}

		Counting(operationsSpy, operationsSpy)

		want := []string{
			writeOperationSignature,
			sleepOperationSignature,
			writeOperationSignature,
			sleepOperationSignature,
			writeOperationSignature,
			sleepOperationSignature,
			writeOperationSignature,
		}

		if !reflect.DeepEqual(want, operationsSpy.Operations) {
			t.Errorf("wanted calls %v got %v", want, operationsSpy.Operations)
		}
	})
}
