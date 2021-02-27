package main

import "testing"

func TestRepeat(t *testing.T) {
	repeticoes := Repeat("a")
	esperado := "aaaaa"

	if repeticoes != esperado {
		t.Errorf("esperado '%s' mas obteve '%s'", esperado, repeticoes)
	}
}
