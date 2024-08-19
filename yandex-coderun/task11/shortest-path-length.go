package main

import "fmt"

func BFS(graph [][]int, start, end int) int {
	n := len(graph)
	visited := make([]bool, n)
	queue := []int{start}
	distances := make([]int, n)
	for i := range distances {
		distances[i] = -1
	}
	distances[start] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			return distances[current]
		}

		for neighbor := 0; neighbor < n; neighbor++ {
			if graph[current][neighbor] == 1 && !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
				distances[neighbor] = distances[current] + 1
			}
		}
	}

	return -1
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
	start--
	end--

	result := BFS(graph, start, end)
	fmt.Println(result)
}
