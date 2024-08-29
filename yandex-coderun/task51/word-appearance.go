package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var text string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		text += line + " "
	}

	words := strings.Fields(text)
	wordCount := make(map[string]int)
	results := []int{}

	for _, word := range words {
		results = append(results, wordCount[word])
		wordCount[word]++
	}

	for _, count := range results {
		fmt.Print(count, " ")
	}

	fmt.Println()
}
