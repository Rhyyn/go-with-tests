package main

import (
	"fmt"
)

const (
	french             = "French"
	spanish            = "Spanish"
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

const englishHelloPrefix = "Hello, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPreifx(language) + name
}

func greetingPreifx(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
