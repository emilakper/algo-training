package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	maze := make([][]byte, n)
	for i := 0; i < n; i++ {
		var row string
		fmt.Scan(&row)
		maze[i] = []byte(row)
	}

	var startRow, startCol int
	fmt.Scan(&startRow, &startCol)

	startRow--
	startCol--

	fmt.Println(calculateRoomArea(maze, startRow, startCol))
}

func calculateRoomArea(maze [][]byte, startRow, startCol int) int {
	n := len(maze)
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
	}

	return dfs(maze, visited, startRow, startCol)
}

func dfs(maze [][]byte, visited [][]bool, row, col int) int {
	if row < 0 || row >= len(maze) || col < 0 || col >= len(maze) {
		return 0
	}
	if visited[row][col] || maze[row][col] == '*' {
		return 0
	}

	visited[row][col] = true
	area := 1

	area += dfs(maze, visited, row-1, col)
	area += dfs(maze, visited, row+1, col)
	area += dfs(maze, visited, row, col-1)
	area += dfs(maze, visited, row, col+1)

	return area
}
