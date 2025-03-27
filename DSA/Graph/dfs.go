package graph

import (
	"fmt"
)

type Graphs struct {
	adjList map[int][]int
	visited map[int]bool
}

func NewGraphs() *Graphs {
	return &Graphs{
		adjList: make(map[int][]int),
		visited: make(map[int]bool),
	}
}

func (g *Graphs) AddEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
}

func (g *Graphs) DFS(node int) {
	if g.visited[node] {
		return
	}

	g.visited[node] = true
	fmt.Printf("%d ", node)

	for _, neighbor := range g.adjList[node] {
		if !g.visited[neighbor] {
			g.DFS(neighbor)
		}
	}
}
func Dfs() {
	graph := NewGraphs()

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 5)

	fmt.Print("DFS Traversal: ")
	graph.DFS(0)
}  