package main

import "fmt"

const russian = "Russian"
const french = "French"
const englishHelloPrefix = "Hello, "
const russianHelloPrefix = "Привет, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}
	if lang == russian {
		return russianHelloPrefix + name + "!"
	}
	if lang == french {
		return frenchHelloPrefix + name + "!"
	}
	return englishHelloPrefix + name + "!"
}

func main() {
	fmt.Println(Hello("world", ""))
}
