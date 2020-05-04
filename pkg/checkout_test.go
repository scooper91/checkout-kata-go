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

var singleItems = []struct {
	item string
	value int
} {
	{"A", 50},
	{"B", 30},
	{"C", 20},
	{"D", 15},
}

func TestSingleItemReturnsCorrectPrice(t *testing.T) {
	for _, singleItem := range singleItems {
		t.Run(singleItem.item, func(t *testing.T) {
			total := checkout(singleItem.item)
			AssertTotalPrice(t, total, singleItem.value)
		})
	}
}

var multipleItems = []struct {
	items string
	total int
} {
	{"ABCD", 115},
	{"ABADCCD", 200},
	{"CCC", 60},
}
func TestMultipleItemsReturnsTotal(t *testing.T) {
	for _, multipleItems := range multipleItems {
		t.Run(multipleItems.items, func(t *testing.T) {
			total := checkout(multipleItems.items)
			AssertTotalPrice(t, total, multipleItems.total)
		})
	}
}

func Test3AsReturnsDiscountedPrice(t *testing.T) {
	total := checkout("AAA")
	AssertTotalPrice(t, total, 130)
}

func Test3AsWithOtherItemsReturnsDiscountedPrice(t *testing.T) {
	total := checkout("CADABA")
	AssertTotalPrice(t, total, 195)
}

func Test4AsReturnsDiscountedPrice(t *testing.T) {
	total := checkout("AAAA")
	AssertTotalPrice(t, total, 180)
}

func Test5AsReturnsDiscountedPrice(t *testing.T) {
	total := checkout("AAAAA")
	AssertTotalPrice(t, total, 230)
}
