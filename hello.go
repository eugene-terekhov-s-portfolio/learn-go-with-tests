package main

import "fmt"

const russian = "Russian"
const englishHelloPrefix = "Hello, "
const russianHelloPrefix = "Привет, "

func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}
	if lang == russian {
		return russianHelloPrefix + name + "!"
	}
	return englishHelloPrefix + name + "!"
}

func main() {
	fmt.Println(Hello("world", ""))
}
