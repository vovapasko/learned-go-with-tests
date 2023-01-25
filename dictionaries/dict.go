package dictionaries

type Dictionary map[string]string

const (
	notFoundErrorValue     = DictionaryError("could not find the word you were looking for")
	wordAlreadyExistsValue = DictionaryError("word already exists")
)

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

func (dictionary Dictionary) Search(key string) (string, error) {
	word, ok := dictionary[key]
	if !ok {
		return "", notFoundErrorValue
	}
	return word, nil
}

func (dictionary Dictionary) Add(key, word string) error {
	existingWord, _ := dictionary.Search(key)
	if existingWord != "" {
		return wordAlreadyExistsValue
	}
	dictionary[key] = word
	return nil
}

func (dictionary Dictionary) Update(key, word string) {
	dictionary[key] = word
}
