package main

import "fmt"

func Sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func SumSlice(numbersToSum ...[]int) []int {
	var sum []int
	for _, numbers := range numbersToSum {
		sum = append(sum, Sum(numbers))
	}

	return sum
}

func SumTotaFinalSlice(numbersToSum ...[]int) []int {
	var sum []int
	for _, numbers := range numbersToSum {
		final := numbers[1:]
		sum = append(sum, Sum(final))
	}

	return sum
}

func main() {
	fmt.Println("Sum slices")
}
