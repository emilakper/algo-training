package main

import (
	"fmt"
	"sort"
)

func maxIndex(data []int) int {
	m := -1 << 31
	maxIndex := 0
	for i, value := range data {
		if value > m {
			m = value
			maxIndex = i
		}
	}
	return maxIndex
}

func place(data []int, member int) int {
	ans := 0
	for _, num := range data {
		if num > member {
			ans++
		}
	}
	return ans + 1
}

func main() {
	var n int
	fmt.Scan(&n)
	data := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	k := maxIndex(data)
	var candidates []int
	for i := k + 1; i < len(data)-1; i++ {
		if data[i]%10 == 5 && data[i+1] < data[i] {
			candidates = append(candidates, data[i])
		}
	}

	if len(candidates) == 0 {
		fmt.Println(0)
	} else {
		sort.Ints(candidates)
		fmt.Println(place(data, candidates[len(candidates)-1]))
	}
}
