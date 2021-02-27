package main

import (
	"reflect"
	"testing"
)

func TestSumSlice(t *testing.T) {
	t.Run("Sum total of slice ", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Sum(numbers)
		expected := 15

		if result != expected {
			t.Errorf("expeted '%d' but result is '%d'. Numbers = %v", expected, result, numbers)
		}
	})

	t.Run("Add two slices and create a new slice with the sum of the two slices", func(t *testing.T) {
		result := SumSlice([]int{1, 3}, []int{2, 4})
		expected := []int{4, 6}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expeted '%v' but result is '%v'.", expected, result)
		}
	})

	t.Run("calculates the totals of all 'endings' of each slice", func(t *testing.T) {
		result := SumTotaFinalSlice([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expeted '%v' but result is '%v'.", expected, result)
		}
	})
}
