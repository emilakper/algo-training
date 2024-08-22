package main

import (
	"container/list"
	"fmt"
)

type Point struct {
	x, y, z int
}

func bfs(cave [][][]rune, start Point, N int) int {
	dx := []int{1, -1, 0, 0, 0, 0}
	dy := []int{0, 0, 1, -1, 0, 0}
	dz := []int{0, 0, 0, 0, 1, -1}

	queue := list.New()
	queue.PushBack(start)
	visited := make(map[Point]bool)
	visited[start] = true
	distance := make(map[Point]int)
	distance[start] = 0

	for queue.Len() > 0 {
		curr := queue.Remove(queue.Front()).(Point)
		if curr.z == 0 {
			return distance[curr]
		}

		for i := 0; i < 6; i++ {
			nx, ny, nz := curr.x+dx[i], curr.y+dy[i], curr.z+dz[i]
			if nx >= 0 && nx < N && ny >= 0 && ny < N && nz >= 0 && nz < N {
				next := Point{nx, ny, nz}
				if cave[nz][ny][nx] != '#' && !visited[next] {
					queue.PushBack(next)
					visited[next] = true
					distance[next] = distance[curr] + 1
				}
			}
		}
	}
	return -1
}

func main() {
	var N int
	fmt.Scan(&N)

	cave := make([][][]rune, N)
	var start Point

	for z := 0; z < N; z++ {
		cave[z] = make([][]rune, N)
		for y := 0; y < N; y++ {
			cave[z][y] = make([]rune, N)
			var line string
			fmt.Scan(&line)
			for x, char := range line {
				cave[z][y][x] = char
				if char == 'S' {
					start = Point{x, y, z}
				}
			}
		}
	}

	minDistance := bfs(cave, start, N)
	fmt.Println(minDistance)
}
