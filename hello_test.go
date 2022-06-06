package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	assetCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Avery", "EN")
		want := "Hello, Avery"
		assetCorrectMessage(t, got, want)
	})

	t.Run("hello to world if no name", func(t *testing.T) {
		got := Hello("", "EN")
		want := "Hello, world"
		assetCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Avery", "ES")
		want := "Hola, Avery"
		assetCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Avery", "FR")
		want := "Bonjour, Avery"
		assetCorrectMessage(t, got, want)
	})
}
