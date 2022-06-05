package main

import "testing"

func TestHello(t *testing.T) {
	assetCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Avery")
		want := "Hello, Avery"
		assetCorrectMessage(t, got, want)
	})

	t.Run("hello to world if no name", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		assetCorrectMessage(t, got, want)
	})
}
