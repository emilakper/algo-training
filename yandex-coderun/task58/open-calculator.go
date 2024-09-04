package main

import (
	"fmt"
)

func main() {
	var x, y, z int
	var N string

	fmt.Scan(&x, &y, &z)
	fmt.Scan(&N)

	availableButtons := map[rune]bool{
		rune('0' + x): true,
		rune('0' + y): true,
		rune('0' + z): true,
	}

	neededButtons := 0
	for _, digit := range N {
		if !availableButtons[digit] {
			neededButtons++
			availableButtons[digit] = true
		}
	}

	fmt.Println(neededButtons)
}
