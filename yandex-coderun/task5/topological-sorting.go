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
	edges     map[int][]int
}

func NewGraph(verticies int) *Graph {
	return &Graph{
		verticies: verticies,
		edges:     make(map[int][]int),
	}
}

func (g *Graph) AddEdge(u, v int) {
	g.edges[u] = append(g.edges[u], v)
}

func (g *Graph) TopologicalSort() ([]int, error) {
	inDegree := make([]int, g.verticies+1)

	for u := range g.edges {
		for _, v := range g.edges[u] {
			inDegree[v]++
		}
	}

	queue := []int{}
	for i := 1; i <= g.verticies; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	topologicalOrder := []int{}

	visited := 0

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		topologicalOrder = append(topologicalOrder, u)
		visited++

		for _, v := range g.edges[u] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	if visited != g.verticies {
		return nil, fmt.Errorf("Graph has cycles")
	}

	return topologicalOrder, nil
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

	result, err := g.TopologicalSort()

	if err != nil {
		fmt.Println("-1")
		return
	}

	for _, v := range result {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
