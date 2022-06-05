package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Avery")
	want := "Hello, Avery"

	if got != want {
		t.Errorf("got %q  want %q", got, want)
	}
}
