package main

import "testing"

func TestRepeat(t *testing.T) {
	result := Repeat("a")
	expected := "aaaaa"

	if result != expected {
		t.Errorf("expected '%s' but result is '%s'", expected, result)
	}
}
