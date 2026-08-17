package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("Hello specific user", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertTestMessage(t, got, want)
	})

	t.Run("Hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertTestMessage(t, got, want)
	})

	t.Run("Hello in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertTestMessage(t, got, want)
	})

	t.Run("Hello in French", func(t *testing.T) {
		got := Hello("Chloé", "French")
		want := "Bonjour, Chloé"

		assertTestMessage(t, got, want)
	})

}

func assertTestMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
