package main

import "fmt"

func Hello(name string, language string) string {
	const spanishHelloPrefix = "Hola, "
	const englishHelloPrefix = "Hello, "
	languageGreetingPrefix := englishHelloPrefix
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		languageGreetingPrefix = spanishHelloPrefix
	}
	return languageGreetingPrefix + name
}

func main() {
	fmt.Println(Hello("world", "Spanish"))
}
