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
	for _, itemTotals := range multipleItems {
		t.Run(itemTotals.items, func(t *testing.T) {
			total := checkout(itemTotals.items)
			AssertTotalPrice(t, total, itemTotals.total)
		})
	}
}

var singleADiscount = []struct {
	items string
	total int
} {
	{"AAA", 130},
	{"AAAA", 180},
	{"AAAAA", 230},
	{"CADABA", 195},
}
func TestBasketsWithADiscountReturnsDiscountedPrice(t *testing.T) {
	for _, itemTotals := range singleADiscount {
		t.Run(itemTotals.items, func(t *testing.T) {
			total := checkout(itemTotals.items)
			AssertTotalPrice(t, total, itemTotals.total)
		})
	}
}

var multipleADiscounts = []struct {
	items string
	total int
} {
	{"AAAAAA", 260},
	{"AAAAAAAAA", 390},
	{"AAAAAAAAAA", 440},
	{"AACAABAA", 310},
}
func TestBasketsWithMultipleADiscountsReturnsDiscountedPrice(t *testing.T) {
	for _, itemTotals := range multipleADiscounts {
		t.Run(itemTotals.items, func(t *testing.T) {
			total := checkout(itemTotals.items)
			AssertTotalPrice(t, total, itemTotals.total)
		})
	}
}

var singleBDiscount = []struct {
	items string
	total int
} {
	{"BB", 45},
	{"BBB", 75},
	{"BABCD", 130},
}
func TestBasketsWithBDiscountReturnsDiscountedPrice(t *testing.T) {
	for _, itemTotals := range singleBDiscount {
		t.Run(itemTotals.items, func(t *testing.T) {
			total := checkout(itemTotals.items)
			AssertTotalPrice(t, total, itemTotals.total)
		})
	}
}

func TestBasketWithAAndBDiscountsReturnsDiscountedPrice(t *testing.T) {
	total := checkout("AAABB")
	AssertTotalPrice(t, total, 175)
}
