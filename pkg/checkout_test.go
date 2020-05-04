package main

import "testing"

func TestEmptyBasketReturns0(t *testing.T) {
	const expected = 0
	total := checkout("")
	if total != expected {
		t.Errorf("Expected %d; Got %d", expected, total)
	}
}

func TestABasketReturns50(t *testing.T) {
	const expected = 50
	total := checkout("A")
	if total != expected {
		t.Errorf("Expected %d; Got %d", expected, total)
	}
}
