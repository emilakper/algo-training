package main

import (
	"bufio"
	"fmt"
	"os"
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

func (g *Graph) isBipartiteDFS(node int, color []int, c int) bool {
	if color[node] != 0 {
		return color[node] == c
	}
	color[node] = c
	for _, neighbor := range g.adjList[node] {
		if !g.isBipartiteDFS(neighbor, color, -c) {
			return false
		}
	}
	return true
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

	color := make([]int, n+1)
	isBipartite := true

	for i := 1; i <= n; i++ {
		if color[i] == 0 {
			if !g.isBipartiteDFS(i, color, 1) {
				isBipartite = false
				break
			}
		}
	}

	if isBipartite {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
