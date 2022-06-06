package main

import (
	"go_with_tests/integers"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := integers.Add(6, 9)
	expected := 15

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
