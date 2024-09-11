package main

import (
	"fmt"
)

func isSymmetric(seq []int) bool {
	n := len(seq)
	for i := 0; i < n/2; i++ {
		if seq[i] != seq[n-i-1] {
			return false
		}
	}
	return true
}

func makeSymmetric(seq []int) (int, []int) {
	n := len(seq)
	for i := 0; i < n; i++ {
		if isSymmetric(seq[i:]) {
			additional := seq[:i]
			for j := 0; j < len(additional)/2; j++ {
				additional[j], additional[len(additional)-j-1] = additional[len(additional)-j-1], additional[j]
			}
			return len(additional), additional
		}
	}
	return 0, nil
}

func main() {
	var n int
	fmt.Scan(&n)

	seq := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&seq[i])
	}

	if isSymmetric(seq) {
		fmt.Println(0)
	} else {
		count, additional := makeSymmetric(seq)
		fmt.Println(count)
		for _, num := range additional {
			fmt.Print(num, " ")
		}
		fmt.Println()
	}
}
