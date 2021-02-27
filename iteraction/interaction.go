package main

import "fmt"

func Repeat(letter string) string {
	var TotalCharacters string
	for i := 0; i < 10; i++ {
		TotalCharacters += letter
	}
	return TotalCharacters
}

func main() {
	fmt.Println("Test the Interaction")
}
