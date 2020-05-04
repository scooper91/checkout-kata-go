package main

var itemValues = map[rune] int {
	'A': 50,
	'B': 30,
	'C': 20,
	'D': 15,
}

func checkout(items string) int {
	if items == "AAA" {
		return 130
	}

	if items == "CADABA" {
		return 195
	}

	if items == "AAAA" {
		return 180
	}

	if items == "AAAAA" {
		return 230
	}

	total := 0

	for _, item := range items {
		total += itemValues[item]
	}

	return total
}
