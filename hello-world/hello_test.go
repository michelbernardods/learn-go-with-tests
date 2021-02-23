package main

import "testing"

func TestOla(t *testing.T) {

	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Michel", "")
		esperado := "Hello Michel"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz 'Olá, mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Hello world"

		verificaMensagemCorreta(t, resultado, esperado)

	})

	t.Run("in english", func(t *testing.T) {
		resultado := Ola("Michel", "english")
		esperado := "Hello Michel"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("in espanhol", func(t *testing.T) {
		resultado := Ola("Michel", "espanhol")
		esperado := "Hola Michel"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("in frances", func(t *testing.T) {
		resultado := Ola("Michel", "frances")
		esperado := "Bonjour Michel"

		verificaMensagemCorreta(t, resultado, esperado)
	})
}
