package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dicionary{"test": "this is just a test"}

	t.Run("familiar word", func(t *testing.T) {
		result, _ := dictionary.Search("teste")
		expected := "this is just a test"

		comparaStrings(t, result, expected)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("an error is expected to be fetched.")
		}
	})

	t.Run("unknown word", func(t *testing.T) {
		_, result := dictionary.Search("unknown")

		comparaErro(t, result, ErrNotFound)
	})
}



func comparaStrings(t *testing.T, result, expected string) {
	t.Helper()

	if result != expected {
		t.Errorf("result '%s', expected '%s', dado '%s'", result, expected, "teste")
	}
}

func comparaErro(t *testing.T, result, expected error) {
	t.Helper()

	if result != expected {
		t.Errorf("result erro '%s', expected '%s'", result, expected)
	}
}

func TestAdds(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dicionary{}
		word := "teste"
		definicao := "isso é apenas um teste"

		err := dictionary.Adds(word, definicao)

		comparaErro(t, err, nil)
		comparaDefinicao(t, dictionary, word, definicao)
	})

	t.Run("word existente", func(t *testing.T) {
		word := "teste"
		definicao := "isso é apenas um teste"
		dictionary := Dicionary{word: definicao}
		err := dictionary.Adds(word, "teste novo")

		comparaErro(t, err, ErrWordNonexistent)
		comparaDefinicao(t, dictionary, word, definicao)
	})
}

func comparaDefinicao(t *testing.T, dictionary Dicionary, word, definicao string) {
	t.Helper()

	result, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("deveria ter encontrado word adicionada:", err)
	}

	if definicao != result {
		t.Errorf("result '%s',  expected '%s'", result, definicao)
	}
}

func TestUpdate(t *testing.T) {
	t.Run("word existente", func(t *testing.T) {
		word := "teste"
		definicao := "isso é apenas um teste"
		novaDefinicao := "nova definição"
		dictionary := Dicionary{word: definicao}
		err := dictionary.Update(word, novaDefinicao)

		comparaErro(t, err, nil)
		comparaDefinicao(t, dictionary, word, novaDefinicao)
	})

	t.Run("word nova", func(t *testing.T) {
		word := "teste"
		definicao := "isso é apenas um teste"
		dictionary := Dicionary{}

		err := dictionary.Update(word, definicao)

		comparaErro(t, err, ErrWordNonexistent)
	})
}

func TestDelete(t *testing.T) {
	word := "teste"
	dictionary := Dicionary{word: "definição de teste"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("espera-se que '%s' seja deletado", word)
	}
}