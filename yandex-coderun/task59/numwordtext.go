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
		text += scanner.Text() + " "
	}

	words := strings.Fields(text)
	uniqueWords := make(map[string]bool)

	for _, word := range words {
		uniqueWords[word] = true
	}

	fmt.Println(len(uniqueWords))
}
