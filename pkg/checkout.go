package main

var itemValues = map[string] int {
	"A": 50,
	"B": 30,
	"C": 20,
}

func checkout(items string) int {
	price, itemExists := itemValues[items]

	if itemExists {
		return price
	}

	return 0
}
