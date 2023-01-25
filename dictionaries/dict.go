package dictionaries

type Dictionary map[string]string

const (
	wordNotFoundError      = DictionaryError("could not find the word you were looking for")
	wordAlreadyExistsError = DictionaryError("word already exists")
	wordDoesNotExistError  = DictionaryError("word does not exist")
)

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

func (dictionary Dictionary) Search(key string) (string, error) {
	word, ok := dictionary[key]
	if !ok {
		return "", wordNotFoundError
	}
	return word, nil
}

func (dictionary Dictionary) Add(key, word string) error {
	existingWord, _ := dictionary.Search(key)
	if existingWord != "" {
		return wordAlreadyExistsError
	}
	dictionary[key] = word
	return nil
}

func (dictionary Dictionary) Update(key, word string) {
	dictionary[key] = word
}

func (dictionary Dictionary) Delete(key string) error {
	_, err := dictionary.Search(key)
	if err == wordNotFoundError {
		return wordDoesNotExistError
	}
	delete(dictionary, key)
	return nil
}
