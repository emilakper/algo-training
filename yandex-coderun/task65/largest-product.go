package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	nums := strings.Split(input, " ")

	arr := make([]int, len(nums))
	for i, numStr := range nums {
		num, _ := strconv.Atoi(numStr)
		arr[i] = num
	}

	sort.Ints(arr)

	n := len(arr)
	max1, max2, max3 := arr[n-1], arr[n-2], arr[n-3]
	min1, min2 := arr[0], arr[1]

	product1 := max1 * max2 * max3
	product2 := max1 * min1 * min2

	if product1 > product2 {
		fmt.Println(max1, max2, max3)
	} else {
		fmt.Println(max1, min1, min2)
	}
}
