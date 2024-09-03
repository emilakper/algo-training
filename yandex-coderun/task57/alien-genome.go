package main

import (
	"fmt"
)

func main() {
	var s1, s2 string
	fmt.Scanln(&s1)
	fmt.Scanln(&s2)

	pairs := make(map[string]bool)
	for i := 0; i < len(s2)-1; i++ {
		pairs[s2[i:i+2]] = true
	}

	ans := 0
	for i := 0; i < len(s1)-1; i++ {
		if pairs[s1[i:i+2]] {
			ans++
		}
	}

	fmt.Println(ans)
}
