package main

func checkout(items string) int {
	if items == "A" {
		return 50
	}

	if items == "B" {
		return 30
	}

	return 0
}
