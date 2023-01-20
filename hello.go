package main

import "fmt"

func Hello(name string) string {
	const englishHelloPrefix = "Hello, "
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
