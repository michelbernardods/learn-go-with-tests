package main

import "fmt"

const ola = "Ola "
const hello = "Hello "
const hola = "Hola "
const bonjour = "Bonjour "

const portugues = "portugues"
const english = "english"
const espanhol = "espanhol"
const frances = "frances"

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return prefixSalutation(language) + name
}

func prefixSalutation(language string) (prefix string) {
	switch language {
	case portugues:
		prefix = ola
	case espanhol:
		prefix = hola
	case frances:
		prefix = bonjour
	default:
		prefix = hello
	}
	return
}

func main() {
	fmt.Println("Hello")
	fmt.Println("Ola")
	fmt.Println("Helo")
	fmt.Println("Bonjour")
}
