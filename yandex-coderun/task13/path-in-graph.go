package main

import (
	"fmt"
)

func bfs(graph [][]int, start, end int) (int, []int) {
	n := len(graph)
	visited := make([]bool, n)
	queue := []int{start}
	visited[start] = true
	prev := make([]int, n)
	for i := range prev {
		prev[i] = -1
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == end {
			break
		}

		for neighbor, isConnected := range graph[node] {
			if isConnected == 1 && !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
				prev[neighbor] = node
			}
		}
	}

	if prev[end] == -1 {
		return -1, nil
	}

	path := []int{}
	for at := end; at != -1; at = prev[at] {
		path = append([]int{at}, path...)
	}

	return len(path) - 1, path
}

func main() {
	var n int
	fmt.Scan(&n)

	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&graph[i][j])
		}
	}

	var start, end int
	fmt.Scan(&start, &end)

	if start == end {
		fmt.Println(0)
		return
	}

	length, path := bfs(graph, start-1, end-1)
	if length == -1 {
		fmt.Println(-1)
	} else {
		fmt.Println(length)
		for _, node := range path {
			fmt.Print(node+1, " ")
		}
		fmt.Println()
	}
}
