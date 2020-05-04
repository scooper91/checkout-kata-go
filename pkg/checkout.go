package main

var itemValues = map[rune] int {
	'A': 50,
	'B': 30,
	'C': 20,
	'D': 15,
}

func checkout(items string) int {
	var itemCounts = map[rune] int {}

	total := 0

	for _, item := range items {
		total += itemValues[item]

		itemCounts[item] = itemCounts[item] + 1
	}

	total -= (itemCounts['A'] / 3) * 20
	total -= (itemCounts['B'] / 2) * 15

	return total
}
