package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	numStrs := strings.Split(input, " ")
	var numbers []int
	for _, numStr := range numStrs {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Ошибка при преобразовании строки в число:", err)
			return
		}
		numbers = append(numbers, num)
	}

	if isMonotonicIncreasing(numbers) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func isMonotonicIncreasing(numbers []int) bool {
	for i := 1; i < len(numbers); i++ {
		if numbers[i] <= numbers[i-1] {
			return false
		}
	}
	return true
}
