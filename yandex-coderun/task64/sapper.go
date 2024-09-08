package main

import (
	"fmt"
)

func main() {
	var N, M, K int
	fmt.Scan(&N, &M, &K)

	field := make([][]int, N)
	for i := range field {
		field[i] = make([]int, M)
	}

	for i := 0; i < K; i++ {
		var p, q int
		fmt.Scan(&p, &q)
		p--
		q--
		field[p][q] = -1
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if field[i][j] == -1 {
				continue
			}
			count := 0
			for di := -1; di <= 1; di++ {
				for dj := -1; dj <= 1; dj++ {
					if di == 0 && dj == 0 {
						continue
					}
					ni, nj := i+di, j+dj
					if ni >= 0 && ni < N && nj >= 0 && nj < M && field[ni][nj] == -1 {
						count++
					}
				}
			}
			field[i][j] = count
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if field[i][j] == -1 {
				fmt.Print("* ")
			} else {
				fmt.Print(field[i][j], " ")
			}
		}
		fmt.Println()
	}
}
