package dictionaries

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"key": "some test value"}
	want := "some test value"
	got := Search(dictionary, "key")
	if want != got {
		t.Errorf("got %s want %s", got, want)
	}
}
