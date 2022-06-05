package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, lang string) string {
	var prefixUsed string = englishHelloPrefix

	if lang == "ES" {
		prefixUsed = spanishHelloPrefix
	}

	if name == "" {
		name = "world"
	}
	return prefixUsed + name
}

func main() {
	fmt.Println(Hello("Avery", "EN"))
}
