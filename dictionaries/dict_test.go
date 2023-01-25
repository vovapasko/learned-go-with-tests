package dictionaries

import "testing"

func TestSearch(t *testing.T) {
	assertStrings := func(got, want string, t testing.TB) {
		if want != got {
			t.Errorf("got %s want %s", got, want)
		}
	}
	assertDefinition := func(dictionary Dictionary, key, word string, t testing.TB) {
		got, err := dictionary.Search(key)
		want := word
		if err != nil {
			t.Fatal("Got error but should nil")
		}
		if got != want {
			t.Errorf("Should find the word")
		}
	}
	assertError := func(err error, t testing.TB) {
		if err == nil {
			t.Fatal("Expected error but didn't get one")
		}
	}
	dictionary := Dictionary{"key": "some test value"}
	t.Run("find existing word", func(t *testing.T) {
		want := "some test value"
		got, _ := dictionary.Search("key")
		assertStrings(got, want, t)
	})
	t.Run("find non existing word", func(t *testing.T) {
		want := notFoundErrorValue.Error()
		_, err := dictionary.Search("non existing key")
		assertError(err, t)
		assertStrings(err.Error(), want, t)
	})
	keyToAdd := "test"
	wordToAdd := "just another word"
	t.Run("add the word", func(t *testing.T) {
		_ = dictionary.Add(keyToAdd, wordToAdd)
		assertDefinition(dictionary, keyToAdd, wordToAdd, t)
	})
	t.Run("add an existing word", func(t *testing.T) {
		err := dictionary.Add(keyToAdd, wordToAdd)
		assertError(err, t)
		assertDefinition(dictionary, keyToAdd, wordToAdd, t)
	})
	t.Run("update an existing word", func(t *testing.T) {
		keyToUpdate := keyToAdd
		wordToUpdate := "new definition"
		dictionary.Update(keyToUpdate, wordToUpdate)
		assertDefinition(dictionary, keyToUpdate, wordToUpdate, t)
	})

}
