package main

import "fmt"

const spanish = "ES"
const french = "FR"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, lang string) string {
	prefix := englishHelloPrefix
	switch lang {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}
	if name == "" {
		name = "world"
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("Avery", "EN"))
}
