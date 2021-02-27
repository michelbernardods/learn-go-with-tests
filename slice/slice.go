package main

import "fmt"

func Sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func SumSlice(numbersToSum ...[]int) (sum []int) {
	numberOfNumbers := len(numbersToSum)
	sum = make([]int, numberOfNumbers)

	for i, number := range numbersToSum {
		sum[i] = Sum(number)
	}

	return
}

func main() {
	fmt.Println("Sum slices")
}
