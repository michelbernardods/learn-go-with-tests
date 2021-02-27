package main

import "fmt"

func Repeat(characters string) string {
	var TotalCharacters string
	for i := 0; i < 5; i++ {
		TotalCharacters += characters
	}
	return TotalCharacters
}

func main() {
	fmt.Println("Test the Interaction")
}
