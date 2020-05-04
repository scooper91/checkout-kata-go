package main

import "testing"

func TestEmptyBasketReturns0(t *testing.T) {
	total := checkout("")
	if total != 0 {
		t.Errorf("Expected 0; Got %d", total)
	}
}
