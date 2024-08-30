package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var n int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)

	synonyms := make(map[string]string)

	for i := 0; i < n; i++ {
		var word1, word2 string
		scanner.Scan()
		word1 = scanner.Text()
		scanner.Scan()
		word2 = scanner.Text()
		synonyms[word1] = word2
		synonyms[word2] = word1
	}

	scanner.Scan()
	query := scanner.Text()

	if synonym, found := synonyms[query]; found {
		fmt.Println(synonym)
	} else {
		fmt.Println("Synonym not found")
	}
}
