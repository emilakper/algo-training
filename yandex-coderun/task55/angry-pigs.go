package main

import (
	"fmt"
	"sort"
)

type Coord struct {
	x, y int
}

func main() {
	var n int
	fmt.Scan(&n)

	birds := make([]Coord, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&birds[i].x, &birds[i].y)
	}

	sort.Slice(birds, func(i, j int) bool {
		if birds[i].x == birds[j].x {
			return birds[i].y < birds[j].y
		}
		return birds[i].x < birds[j].x
	})

	shots := 0
	currentX := -1
	uniqueY := make(map[int]bool)

	for _, bird := range birds {
		if bird.x != currentX {
			shots += len(uniqueY)
			uniqueY = make(map[int]bool)
			currentX = bird.x
		}
		uniqueY[bird.x] = true
	}
	shots += len(uniqueY)

	fmt.Println(shots)
}
