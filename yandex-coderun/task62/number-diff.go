package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader((os.Stdin))
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	numbers := strings.Split(input, " ")
	uniqueNumbers := make(map[int]bool)

	for _, numStr := range numbers {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Ошибка, не число")
			return
		}
		uniqueNumbers[num] = true
	}
	fmt.Println((len(uniqueNumbers)))
}
