package main

import testing "testing"

func TestHello(t *testing.T) {

	t.Run("should say hello to the person", func(t *testing.T) {
		got := hello("Soubhik", "English")
		want := "Hello Soubhik"

		extractingCommon(t, got, want)
	})

	t.Run("should say Hello World when name is not provided", func(t *testing.T) {
		got := hello("", "")
		want := "Hello World"

		extractingCommon(t, got, want)
	})

	t.Run("should say Hello in Spanish if language input is Spanish", func(t *testing.T) {
		got := hello("Steve", "Spanish")
		want := "Hola Steve"

		extractingCommon(t, got, want)
	})

	t.Run("should say Hello in French if language input is French", func(t *testing.T) {
		got := hello("Steve", "French")
		want := "Bonjour Steve"

		extractingCommon(t, got, want)
	})
}

func extractingCommon(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}
