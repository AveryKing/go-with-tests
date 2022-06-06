package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(6, 9)
	expected := 15

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
