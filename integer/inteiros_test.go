package inteiros

import "testing"

func TestAdicionadores(t *testing.T) {
	soma := Adiciona(2, 2)
	resultado := 4

	if soma != resultado {
		t.Errorf("Esperado %d resoltado %d ", resultado, soma)
	}
}
