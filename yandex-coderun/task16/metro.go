package main

import (
	"fmt"
)

type Station struct {
	Number int
	Lines  []int
}

type Line struct {
	Stations []int
}

type Graph struct {
	Stations []Station
	Lines    []Line
}

func buildGraph(N int, lines []Line) Graph {
	stations := make([]Station, N+1)
	for i := 1; i <= N; i++ {
		stations[i].Number = i
	}

	for i, line := range lines {
		for _, station := range line.Stations {
			stations[station].Lines = append(stations[station].Lines, i)
		}
	}

	return Graph{Stations: stations, Lines: lines}
}

func bfs(graph Graph, start, end int) int {
	if start == end {
		return 0
	}

	visited := make([]bool, len(graph.Stations))
	queue := []struct {
		station   int
		transfers int
	}{{start, 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current.station] {
			continue
		}
		visited[current.station] = true

		for _, line := range graph.Stations[current.station].Lines {
			for _, station := range graph.Lines[line].Stations {
				if station == end {
					return current.transfers
				}
				if !visited[station] {
					queue = append(queue, struct {
						station   int
						transfers int
					}{station, current.transfers + 1})
				}
			}
		}
	}

	return -1
}

func main() {
	var N, M int
	fmt.Scan(&N)
	fmt.Scan(&M)

	lines := make([]Line, M)
	for i := 0; i < M; i++ {
		var Pi int
		fmt.Scan(&Pi)
		lines[i].Stations = make([]int, Pi)
		for j := 0; j < Pi; j++ {
			fmt.Scan(&lines[i].Stations[j])
		}
	}

	var A, B int
	fmt.Scan(&A)
	fmt.Scan(&B)

	graph := buildGraph(N, lines)
	result := bfs(graph, A, B)

	fmt.Println(result)
}
