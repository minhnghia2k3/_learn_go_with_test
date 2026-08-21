package main

import (
	"errors"
	"testing"
)

func TestDictionary(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}

	t.Run("search known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("search not exist word", func(t *testing.T) {
		_, got := dict.Search("unknown")
		want := ErrNotFoundWord

		assertError(t, got, want)
	})
}

func TestAdd(t *testing.T) {
	dict := Dictionary{}

	t.Run("new word", func(t *testing.T) {
		word := "test"
		err := dict.Add("test", "this is just a test")
		want := "this is just a test"

		assertError(t, err, nil)
		assertDefinition(t, dict, word, want)
	})

	t.Run("existing word", func(t *testing.T) {
		got := dict.Add("test", "this is just a test")

		assertError(t, got, ErrWordExists)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		dict := Dictionary{word: "this is just a test"}
		newDefinition := "new definition"

		dict.Update(word, newDefinition)

		assertDefinition(t, dict, word, newDefinition)
	})

	t.Run("non-existing word", func(t *testing.T) {
		word := "non-exist"
		definition := "this is just a test"
		dict := Dictionary{}

		err := dict.Update(word, definition)

		assertError(t, err, ErrWordNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete existing word", func(t *testing.T) {
		word := "test"
		dict := Dictionary{word: "this is just a test"}

		dict.Delete(word)

		_, err := dict.Search(word)
		assertError(t, err, ErrNotFoundWord)
	})

	t.Run("delete non-existing word", func(t *testing.T) {
		dict := Dictionary{}

		err := dict.Delete("non-existing")

		assertError(t, err, ErrWordNotExist)
	})
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dict Dictionary, word, want string) {
	t.Helper()
	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should add new word")
	}

	assertString(t, got, want)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if !errors.Is(got, want) {
		t.Errorf("got '%v', want '%v'", got, want)
	}
}
