package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Graph struct {
	verticies int
	adjList   map[int][]int
}

func NewGraph(verticies int) *Graph {
	return &Graph{
		verticies: verticies,
		adjList:   make(map[int][]int),
	}
}

func (g *Graph) AddEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
	g.adjList[v] = append(g.adjList[v], u)
}

func (g *Graph) DFS(start int, visited []bool, component *[]int) {
	visited[start] = true
	*component = append(*component, start)

	for _, neighbor := range g.adjList[start] {
		if !visited[neighbor] {
			g.DFS(neighbor, visited, component)
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	firstLine := scanner.Text()
	parts := strings.Fields(firstLine)

	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	g := NewGraph(n)

	for i := 0; i < m; i++ {
		scanner.Scan()
		edgeLine := scanner.Text()
		edgeParts := strings.Fields(edgeLine)
		u, _ := strconv.Atoi(edgeParts[0])
		v, _ := strconv.Atoi(edgeParts[1])
		g.AddEdge(u, v)
	}

	visited := make([]bool, n+1)
	components := [][]int{}

	for i := 1; i <= n; i++ {
		if !visited[i] {
			component := []int{}
			g.DFS(i, visited, &component)
			components = append(components, component)
		}
	}

	fmt.Println(len(components))

	for _, component := range components {
		sort.Ints(component)
		fmt.Println(len(component))
		for _, vertex := range component {
			fmt.Printf("%d ", vertex)
		}
		fmt.Println()
	}
}
