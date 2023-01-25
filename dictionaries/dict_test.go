package dictionaries

import "testing"

func TestSearch(t *testing.T) {
	assertStrings := func(got, want string, t testing.TB) {
		if want != got {
			t.Errorf("got %s want %s", got, want)
		}
	}
	dictionary := Dictionary{"key": "some test value"}
	t.Run("find existing word", func(t *testing.T) {
		want := "some test value"
		got, _ := dictionary.Search("key")
		assertStrings(got, want, t)
	})
	t.Run("find non existing word", func(t *testing.T) {
		want := notFoundErrorValue
		_, err := dictionary.Search("non existing key")
		if err == nil {
			t.Fatal("Expected error but didn't get one")
		}
		assertStrings(err.Error(), want, t)
	})

}
