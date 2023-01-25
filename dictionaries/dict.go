package dictionaries

import "errors"

type Dictionary map[string]string

var notFoundErrorValue = "could not find the word you were looking for"
var wordAlreadyExistsValue = "word already exists"

func (dictionary Dictionary) Search(key string) (string, error) {
	word, ok := dictionary[key]
	if !ok {
		return "", errors.New(notFoundErrorValue)
	}
	return word, nil
}

func (dictionary Dictionary) Add(key, word string) error {
	existingWord, _ := dictionary.Search(key)
	if existingWord != "" {
		return errors.New(wordAlreadyExistsValue)
	}
	dictionary[key] = word
	return nil
}
