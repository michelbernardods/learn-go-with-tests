package main

import "testing"

func TestHello(t *testing.T) {

	checkCorrectMessage := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		result := Hello("Michel", "")
		expected := "Hello Michel"

		checkCorrectMessage(t, result, expected)
	})

	t.Run("says 'hello world' when an empty string is passed", func(t *testing.T) {
		result := Hello("", "")
		expected := "Hello world"

		checkCorrectMessage(t, result, expected)

	})

	t.Run("in english", func(t *testing.T) {
		result := Hello("Michel", "english")
		expected := "Hello michel"

		checkCorrectMessage(t, result, expected)
	})

	t.Run("in espanhol", func(t *testing.T) {
		result := Hello("Michel", "espanhol")
		expected := "Hola Michel"

		checkCorrectMessage(t, result, expected)
	})

	t.Run("in frances", func(t *testing.T) {
		result := Hello("Michel", "frances")
		expected := "Bonjour Michel"

		checkCorrectMessage(t, result, expected)
	})
}
