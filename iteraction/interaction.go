package main

import "fmt"

func RepeatString(characters string) string {
	var TotalCharacters string
	for i := 0; i < 5; i++ {
		TotalCharacters += characters
	}

	return TotalCharacters
}

func RepeatInteger(integer int) int {
	var TotalInteger = 0
	for i := 0; i < 5; i++ {
		TotalInteger += integer
	}

	return TotalInteger
}

func RepeatFloat(floating float64) float64 {
	var TotalFloat float64
	for i := 0; i < 5; i++ {
		TotalFloat += floating
	}

	return TotalFloat
}

func main() {
	fmt.Println("Test the Interaction")
}
