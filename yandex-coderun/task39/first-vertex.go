package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func dfs(v int, graph map[int][]int, visited []bool) {
	stack := []int{v}
	for len(stack) > 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		visited[v] = true
		for _, u := range graph[v] {
			if !visited[u] {
				stack = append(stack, u)
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	graph := make(map[int][]int)
	reversedGraph := make(map[int][]int)
	for i := 0; i < m; i++ {
		scanner.Scan()
		u, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		graph[u] = append(graph[u], v)
		reversedGraph[v] = append(reversedGraph[v], u)
	}

	visited := make([]bool, n+1)

	dfs(1, reversedGraph, visited)

	var reachable []int
	for i := 1; i <= n; i++ {
		if visited[i] {
			reachable = append(reachable, i)
		}
	}

	sort.Ints(reachable)
	for _, v := range reachable {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
