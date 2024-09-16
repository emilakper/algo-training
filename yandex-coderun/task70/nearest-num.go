package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	var x int
	fmt.Scan(&x)

	closest := array[0]
	minDiff := math.Abs(float64(array[0] - x))

	for _, value := range array {
		diff := math.Abs(float64(value - x))
		if diff < minDiff {
			minDiff = diff
			closest = value
		}
	}

	fmt.Println(closest)
}
