package main

import "testing"

func TestSumSlice(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	result := Sum(numbers)
	expected := 15

	if result != expected {
		t.Errorf("expeted '%d' but result is '%d'. Numbers = %v", expected, result, numbers)
	}
}
