package main

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Test repeat in String", func(t *testing.T) {
		result := RepeatString("a")
		expected := "aaaaa"

		if result != expected {
			t.Errorf("expected '%s' but result is '%s'", expected, result)
		}
	})

	t.Run("Test repeat in Integer", func(t *testing.T) {
		result := RepeatInteger(1)
		expected := 5

		if result != expected {
			t.Errorf("expected '%d' but result is '%d'", expected, result)
		}
	})

	t.Run("Test repeat in Float", func(t *testing.T) {
		result := RepeatFloat(2.5)
		expected := 12.5

		if result != expected {
			t.Errorf("expected '%v' but result is '%v'", expected, result)
		}
	})

}
