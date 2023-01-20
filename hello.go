package main

import "fmt"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	languageGreetingPrefix := getLanguagePrefix(language)
	return languageGreetingPrefix + name
}

func getLanguagePrefix(language string) (languageGreetingPrefix string) {
	const spanishHelloPrefix = "Hola, "
	const frenchHelloPrefix = "Bonjour, "
	const englishHelloPrefix = "Hello, "

	switch language {
	case "French":
		languageGreetingPrefix = frenchHelloPrefix
	case "Spanish":
		languageGreetingPrefix = spanishHelloPrefix
	default:
		languageGreetingPrefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", "Spanish"))
}
