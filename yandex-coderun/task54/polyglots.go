package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var N int
	fmt.Sscanf(scanner.Text(), "%d", &N)

	allLanguages := make(map[string]struct{})
	commonLanguages := make(map[string]struct{})

	for i := 0; i < N; i++ {
		scanner.Scan()
		var M int
		fmt.Sscanf(scanner.Text(), "%d", &M)

		currentLanguages := make(map[string]struct{})
		for j := 0; j < M; j++ {
			scanner.Scan()
			language := scanner.Text()
			currentLanguages[language] = struct{}{}
			allLanguages[language] = struct{}{}
		}

		if i == 0 {
			for lang := range currentLanguages {
				commonLanguages[lang] = struct{}{}
			}
		} else {
			for lang := range commonLanguages {
				if _, exists := currentLanguages[lang]; !exists {
					delete(commonLanguages, lang)
				}
			}
		}
	}

	fmt.Println(len(commonLanguages))
	for lang := range commonLanguages {
		fmt.Println(lang)
	}

	fmt.Println(len(allLanguages))
	for lang := range allLanguages {
		fmt.Println(lang)
	}
}
