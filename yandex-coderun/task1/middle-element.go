package main

import (
	"fmt"
)

func findMiddle(a, b, c int) int {
	if (a >= b && a <= c) || (a >= c && a <= b) {
		return a
	} else if (b >= a && b <= c) || (b >= c && b <= a) {
		return b
	} else {
		return c
	}
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	mid := findMiddle(a, b, c)
	fmt.Println(mid)
}
