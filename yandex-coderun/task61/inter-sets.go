package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseInput(input string) []int {
	parts := strings.Split(input, " ")
	numbers := make([]int, len(parts))
	for i, part := range parts {
		num, _ := strconv.Atoi(part)
		numbers[i] = num
	}
	return numbers
}

func findIntersection(list1, list2 []int) []int {
	elements := make(map[int]bool)
	for _, num := range list1 {
		elements[num] = true
	}

	var intersection []int
	for _, num := range list2 {
		if elements[num] {
			intersection = append(intersection, num)
			delete(elements, num)
		}
	}

	return intersection
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	input1 := scanner.Text()
	scanner.Scan()
	input2 := scanner.Text()

	list1 := parseInput(input1)
	list2 := parseInput(input2)

	intersection := findIntersection(list1, list2)

	sort.Ints(intersection)

	for _, num := range intersection {
		fmt.Print(num, " ")
	}
	fmt.Println()
}
