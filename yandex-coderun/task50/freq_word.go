package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var lines []string

	for {
		line, err := reader.ReadString('\n')
		if err != nil && line == "" {
			break
		}
		lines = append(lines, strings.TrimSpace(line))
	}

	text := strings.Join(lines, " ")
	words := strings.Fields(text)

	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	maxCount := 0
	var maxWords []string
	for word, count := range wordCount {
		if count > maxCount {
			maxCount = count
			maxWords = []string{word}
		} else if count == maxCount {
			maxWords = append(maxWords, word)
		}
	}

	sort.Strings(maxWords)
	fmt.Println(maxWords[0])
}
