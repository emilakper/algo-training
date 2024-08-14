package main

import (
	"fmt"
	"log"
)

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var less, greater []int

	for _, value := range arr[1:] {
		if value <= pivot {
			less = append(less, value)
		} else {
			greater = append(greater, value)
		}
	}

	less = quickSort(less)
	greater = quickSort(greater)

	return append(append(less, pivot), greater...)
}

func main() {
	var arr = make([]int, 3)
	for i := 0; i < len(arr); i++ {
		_, err := fmt.Scan(&arr[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	arr = quickSort(arr)

	fmt.Println(arr[1])
}
