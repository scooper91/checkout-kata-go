package main

func checkout(items string) int {
	if items == "A" {
		return 50
	}

	if items == "B" {
		return 30
	}

	if items == "C" {
		return 20
	}

	return 0
}
