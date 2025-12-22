package test

import "testing"

func TestAddOne(t *testing.T) {
	var (
		a, b, expected int = 2, 3, 6
	)

	if a+b != expected {
		t.Errorf("Expected %d but got %d", expected, a+b)
	}
}
