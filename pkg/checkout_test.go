package main

import "testing"

func AssertTotalPrice(t *testing.T, actual int, expected int) {
	if actual != expected {
		t.Errorf("Expected %d; Got %d", expected, actual)
	}
}

func TestEmptyBasketReturns0(t *testing.T) {
	total := checkout("")
	AssertTotalPrice(t, total, 0)
}

func TestABasketReturns50(t *testing.T) {
	total := checkout("A")
	AssertTotalPrice(t, total, 50)
}

func TestBBasketReturns30(t *testing.T) {
	total := checkout("B")
	AssertTotalPrice(t, total, 30)
}

func TestCBasketReturns20(t *testing.T) {
	total := checkout("C")
	AssertTotalPrice(t, total, 20)
}

func TestDBasketReturns15(t *testing.T) {
	total := checkout("D")
	AssertTotalPrice(t, total, 15)
}
