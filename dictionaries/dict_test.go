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
	assertError := func(gotError error, wantError error, t testing.TB) {
		if gotError == nil {
			t.Fatal("Expected error but didn't get one")
		}
		if gotError != wantError {
			t.Errorf("Got %s but need %s", gotError.Error(), wantError.Error())
		}
	}
	dictionary := Dictionary{"key": "some test value"}
	t.Run("find existing word", func(t *testing.T) {
		want := "some test value"
		got, _ := dictionary.Search("key")
		assertStrings(got, want, t)
	})
	t.Run("find non existing word", func(t *testing.T) {
		want := wordNotFoundError
		_, err := dictionary.Search("non existing key")
		assertError(err, want, t)
	})
	keyToAdd := "test"
	wordToAdd := "just another word"
	t.Run("add the word", func(t *testing.T) {
		_ = dictionary.Add(keyToAdd, wordToAdd)
		assertDefinition(dictionary, keyToAdd, wordToAdd, t)
	})
	t.Run("add an existing word", func(t *testing.T) {
		err := dictionary.Add(keyToAdd, wordToAdd)
		assertError(err, wordAlreadyExistsError, t)
		assertDefinition(dictionary, keyToAdd, wordToAdd, t)
	})
	t.Run("update an existing word", func(t *testing.T) {
		keyToUpdate := keyToAdd
		wordToUpdate := "new definition"
		dictionary.Update(keyToUpdate, wordToUpdate)
		assertDefinition(dictionary, keyToUpdate, wordToUpdate, t)
	})
	t.Run("delete an existing word", func(t *testing.T) {
		keyToDelete := keyToAdd
		err := dictionary.Delete(keyToDelete)
		if err != nil {
			t.Fatal("Got error but shouldn't get one")
		}
	})
	t.Run("delete non existing word", func(t *testing.T) {
		keyToDelete := keyToAdd
		err := dictionary.Delete(keyToDelete)
		assertError(err, wordDoesNotExistError, t)
	})

}
