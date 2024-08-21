package main

import (
	"container/list"
	"fmt"
)

type Point struct {
	x, y int
}

var directions = []Point{
	{-2, -1}, {-1, -2}, {1, -2}, {2, -1},
	{2, 1}, {1, 2}, {-1, 2}, {-2, 1},
}

func isValid(x, y, N, M int) bool {
	return x >= 1 && x <= N && y >= 1 && y <= M
}

func bfs(N, M, S, T int) [][]int {
	dist := make([][]int, N+1)
	for i := range dist {
		dist[i] = make([]int, M+1)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}
	dist[S][T] = 0
	queue := list.New()
	queue.PushBack(Point{S, T})

	for queue.Len() > 0 {
		curr := queue.Remove(queue.Front()).(Point)
		for _, dir := range directions {
			nx, ny := curr.x+dir.x, curr.y+dir.y
			if isValid(nx, ny, N, M) && dist[nx][ny] == -1 {
				dist[nx][ny] = dist[curr.x][curr.y] + 1
				queue.PushBack(Point{nx, ny})
			}
		}
	}

	return dist
}

func main() {
	var N, M, S, T, Q int
	fmt.Scan(&N, &M, &S, &T, &Q)

	dist := bfs(N, M, S, T)

	sum := 0
	impossible := false

	for i := 0; i < Q; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		if dist[x][y] == -1 {
			impossible = true
		}
		sum += dist[x][y]
	}

	if impossible {
		fmt.Println(-1)
	} else {
		fmt.Println(sum)
	}
}
