package main

import "testing"

func SumTestTwoNumbers(t *testing.T) {
	sum := SumTest(2, 2)
	result := 4

	if sum != result {
		t.Errorf("Expected %d result %d ", result, sum)
	}
}
