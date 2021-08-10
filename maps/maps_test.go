package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dicionary{"test": "this is just a test"}

	t.Run("familiar word", func(t *testing.T) {
		result, _ := dictionary.Search("test")
		expected := "this is just a test"

		compareStrings(t, result, expected)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("an error is expected to be fetched.")
		}
	})

	t.Run("unknown word", func(t *testing.T) {
		_, result := dictionary.Search("unknown")

		comparaErr(t, result, ErrNotFound)
	})
}

func compareStrings(t *testing.T, result, expected string) {
	t.Helper()

	if result != expected {
		t.Errorf("result '%s', expected '%s', dado '%s'", result, expected, "test")
	}
}

func comparaErr(t *testing.T, result, expected error) {
	t.Helper()

	if result != expected {
		t.Errorf("result erro '%s', expected '%s'", result, expected)
	}
}

func TestAdds(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dicionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Adds(word, definition)

		comparaErr(t, err, nil)
		comparaDefinicao(t, dictionary, word, definition)
	})

	t.Run("word existing", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dicionary{word: definition}
		err := dictionary.Adds(word, "test novo")

		comparaErr(t, err, ErrWordExist)
		comparaDefinicao(t, dictionary, word, definition)
	})
}

func comparaDefinicao(t *testing.T, dictionary Dicionary, word, definition string) {
	t.Helper()

	result, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should have found word added:", err)
	}

	if definition != result {
		t.Errorf("result '%s',  expected '%s'", result, definition)
	}
}

func TestUpdate(t *testing.T) {
	t.Run("word existente", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		novaDefinicao := "nova definição"
		dictionary := Dicionary{word: definition}
		err := dictionary.Update(word, novaDefinicao)

		comparaErr(t, err, nil)
		comparaDefinicao(t, dictionary, word, novaDefinicao)
	})

	t.Run("word nova", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dicionary{}

		err := dictionary.Update(word, definition)

		comparaErr(t, err, ErrWordNonexistent)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dicionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("expected to be '%s' deleted", word)
	}
}