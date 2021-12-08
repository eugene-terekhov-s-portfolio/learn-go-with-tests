package main

import "fmt"

const russian = "Russian"
const russianHelloPrefix = "Привет, "
const french = "French"
const frenchHelloPrefix = "Bonjour, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "

const englishHelloPrefix = "Hello, "

func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}

	return prefixByLanguage(lang) + name + "!"
}

func prefixByLanguage(lang string) (prefix string) {
	switch lang {
	case russian:
		prefix = russianHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("world", ""))
}
