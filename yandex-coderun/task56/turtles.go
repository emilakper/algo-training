package main

import (
	"fmt"
)

func maxTruthfulTurtles(N int, statements [][2]int) int {
	count := 0
	seen := make(map[[2]int]bool)
	for _, stmt := range statements {
		a, b := stmt[0], stmt[1]
		if a >= 0 && b >= 0 && a+b+1 == N {
			if !seen[[2]int{a, b}] {
				count++
				seen[[2]int{a, b}] = true
			}
		}
	}
	return count
}

func main() {
	var N int
	fmt.Scan(&N)

	statements := make([][2]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&statements[i][0], &statements[i][1])
	}

	result := maxTruthfulTurtles(N, statements)
	fmt.Println(result)
}
