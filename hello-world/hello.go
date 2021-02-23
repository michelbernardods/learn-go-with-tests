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

func Ola(nome, idioma string) string {
	if nome == "" {
		nome = "world"
	}

	return prefixoSaudacao(idioma) + nome
}

func prefixoSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case portugues:
		prefixo = ola
	case espanhol:
		prefixo = hola
	case frances:
		prefixo = bonjour
	default:
		prefixo = hello
	}
	return
}

func main() {
	fmt.Println("Hello")
	fmt.Println("Ola")
	fmt.Println("Helo")
	fmt.Println("Bonjour")

}
