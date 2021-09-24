package main

import "fmt"

const englishLanguage = "English"
const italianLanguage = "Italian"
const frenchLanguage = "French"

const englishHelloPrefix = "Hello, "
const italianHelloPrefix = "Ciao, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case englishLanguage:
		prefix = englishHelloPrefix
	case italianLanguage:
		prefix = italianHelloPrefix
	case frenchLanguage:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
