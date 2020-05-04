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

	for _, item := range items {
		total += itemValues[item]
		if item == 'A' {
			aCount++
		}
	}

	if aCount >= 3 {
		total -= 20
	}

	return total
}
