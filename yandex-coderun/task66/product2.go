package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Scanln(&input)

	nums := strings.Split(input, " ")
	numbers := make([]int, len(nums))

	for i, numStr := range nums {
		num, _ := strconv.Atoi(numStr)
		numbers[i] = num
	}

	max1, max2 := -1000001, -1000001
	min1, min2 := 1000001, 1000001

	for _, num := range numbers {
		if num > max1 {
			max2 = max1
			max1 = num
		} else if num > max2 {
			max2 = num
		}

		if num < min1 {
			min2 = min1
			min1 = num
		} else if num < min2 {
			min2 = num
		}
	}

	product1 := max1 * max2
	product2 := min1 * min2

	if product1 > product2 {
		if max1 > max2 {
			fmt.Println(max2, max1)
		} else {
			fmt.Println(max1, max2)
		}
	} else {
		if min1 < min2 {
			fmt.Println(min1, min2)
		} else {
			fmt.Println(min2, min1)
		}
	}
}
