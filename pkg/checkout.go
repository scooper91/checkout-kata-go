package main

var itemValues = map[rune] int {
	'A': 50,
	'B': 30,
	'C': 20,
	'D': 15,
}

type itemDiscount struct {
	AmountNeeded, Discount int
}

var discounts = map[rune] itemDiscount {
	'A': itemDiscount {3, 20},
	'B': itemDiscount {2, 15},
}

func checkout(items string) int {
	var itemCounts = map[rune] int {}

	total := 0

	for _, item := range items {
		total += itemValues[item]

		itemCounts[item] = itemCounts[item] + 1
	}

	total -= (itemCounts['A'] / discounts['A'].AmountNeeded) * discounts['A'].Discount
	total -= (itemCounts['B'] / discounts['B'].AmountNeeded) * discounts['B'].Discount

	return total
}
