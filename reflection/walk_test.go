package reflection

import "testing"

func TestWalk(t *testing.T) {
	expected := "Chris"
	var got []string

	x := struct {
		Name string
	}{expected}
	Walk(x, func(input string) {
		got = append(got, input)
	})
	if len(got) != 1 {
		t.Errorf("Expected got = %d", len(got))
	}
	if got[0] != expected {
		t.Errorf("Expected %s but got %s", expected, got[0])
	}
}
