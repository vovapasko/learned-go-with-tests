package main

import (
	"log"
	"my-super-project/dependency_injection"
	"net/http"
)

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
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreetHandler)))
}

func MyGreetHandler(writer http.ResponseWriter, request *http.Request) {
	dependency_injection.Greet(writer, "world")
}
