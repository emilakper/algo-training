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

	numbers := strings.Fields(input)
	count := 0

	for i := 1; i < len(numbers)-1; i++ {
		prev, _ := strconv.Atoi(numbers[i-1])
		current, _ := strconv.Atoi(numbers[i])
		next, _ := strconv.Atoi(numbers[i+1])

		if current > prev && current > next {
			count++
		}
	}

	fmt.Println(count)
}
