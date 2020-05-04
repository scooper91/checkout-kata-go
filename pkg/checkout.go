package main

var itemValues = map[rune] int {
	'A': 50,
	'B': 30,
	'C': 20,
	'D': 15,
}

func checkout(items string) int {
	total := 0
	aCount := 0
	bCount := 0

	for _, item := range items {
		total += itemValues[item]
		if item == 'A' {
			aCount++
		}

		if item == 'B' {
			bCount++
		}
	}

	total -= (aCount / 3) * 20
	total -= (bCount / 2) * 15

	return total
}
