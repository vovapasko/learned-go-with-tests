package dependency_injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	err := Greet(&buffer, "Chris")
	if err != nil {
		t.Fatal("Error was raised")
	}
	get := buffer.String()
	want := "Hello, Chris"
	if get != want {
		t.Errorf("Get %s want %s", get, want)
	}
}
