package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func canOrderContainers(priorities []float64) int {
	stack := []float64{}
	expected := 1.0

	for _, priority := range priorities {
		for len(stack) > 0 && stack[len(stack)-1] == expected {
			stack = stack[:len(stack)-1]
			expected += 1.0
		}
		if priority == expected {
			expected += 1.0
		} else {
			stack = append(stack, priority)
		}
	}

	for len(stack) > 0 {
		if stack[len(stack)-1] == expected {
			stack = stack[:len(stack)-1]
			expected += 1.0
		} else {
			return 0
		}
	}

	return 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < N; i++ {
		scanner.Scan()
		input := strings.Fields(scanner.Text())
		K, _ := strconv.Atoi(input[0])
		priorities := make([]float64, K)
		for j := 0; j < K; j++ {
			priorities[j], _ = strconv.ParseFloat(input[j+1], 64)
		}
		result := canOrderContainers(priorities)
		writer.WriteString(fmt.Sprintf("%d\n", result))
	}
}
