package graphtraversal

import "fmt"

type Graph struct {
	Nodes []Node
}

type Node struct {
	Num       int
	Visited   bool
	Marked    bool
	Neighbors []int
}

func GetGraph(n int) Graph {
	var chosenGraph Graph

	graph1a := Graph{
		Nodes: []Node{
			{Num: 1, Visited: false, Marked: false, Neighbors: []int{2, 4}},
			{Num: 2, Visited: false, Marked: false, Neighbors: []int{1, 5, 7}},
			{Num: 3, Visited: false, Marked: false, Neighbors: []int{4, 6, 7}},
			{Num: 4, Visited: false, Marked: false, Neighbors: []int{1, 3, 5}},
			{Num: 5, Visited: false, Marked: false, Neighbors: []int{2, 4, 6, 7}},
			{Num: 6, Visited: false, Marked: false, Neighbors: []int{3, 5}},
			{Num: 7, Visited: false, Marked: false, Neighbors: []int{2, 3, 5}},
		},
	}

	graph1b := Graph{
		Nodes: []Node{
			{Num: 1, Visited: false, Marked: false, Neighbors: []int{4, 5}},
			{Num: 2, Visited: false, Marked: false, Neighbors: []int{3, 4}},
			{Num: 3, Visited: false, Marked: false, Neighbors: []int{2, 4}},
			{Num: 4, Visited: false, Marked: false, Neighbors: []int{1, 2, 3, 7}},
			{Num: 5, Visited: false, Marked: false, Neighbors: []int{1, 6}},
			{Num: 6, Visited: false, Marked: false, Neighbors: []int{5, 7}},
			{Num: 7, Visited: false, Marked: false, Neighbors: []int{4, 6}},
		},
	}

	graph1c := Graph{
		Nodes: []Node{
			{Num: 1, Visited: false, Marked: false, Neighbors: []int{2, 4, 5}},
			{Num: 2, Visited: false, Marked: false, Neighbors: []int{3, 4, 7}},
			{Num: 3, Visited: false, Marked: false, Neighbors: []int{4}},
			{Num: 4, Visited: false, Marked: false, Neighbors: []int{6, 7}},
			{Num: 5, Visited: false, Marked: false, Neighbors: []int{4}},
			{Num: 6, Visited: false, Marked: false, Neighbors: []int{5}},
			{Num: 7, Visited: false, Marked: false, Neighbors: []int{6}},
		},
	}

	graph1d := Graph{
		Nodes: []Node{
			{Num: 1, Visited: false, Marked: false, Neighbors: []int{2, 4, 5}},
			{Num: 2, Visited: false, Marked: false, Neighbors: []int{3}},
			{Num: 3, Visited: false, Marked: false, Neighbors: []int{7}},
			{Num: 4, Visited: false, Marked: false, Neighbors: []int{2, 3, 6, 7}},
			{Num: 5, Visited: false, Marked: false, Neighbors: []int{4}},
			{Num: 6, Visited: false, Marked: false, Neighbors: []int{7}},
			{Num: 7, Visited: false, Marked: false, Neighbors: []int{}},
		},
	}

	switch n {
	case 1:
		chosenGraph = graph1a
	case 2:
		chosenGraph = graph1b
	case 3:
		chosenGraph = graph1c
	case 4:
		chosenGraph = graph1d
	}

	return chosenGraph
}

func Visit(n *Node) {
	n.Visited = true
}

func Mark(n *Node) {
	n.Marked = true
}

func DepthFirstTraversal(g Graph, n *Node) {
	Visit(n)
	Mark(n)
	fmt.Print(n.Num, " ")

	for i := 0; i < len(n.Neighbors); i++ {
		if !g.Nodes[n.Neighbors[i]-1].Marked {
			DepthFirstTraversal(g, &g.Nodes[n.Neighbors[i]-1])
		}
	}
}

var breadthQueue []*Node

func Enqueue(g *Graph, n *Node) {
	breadthQueue = append(breadthQueue, n)
	for i := 0; i < len(n.Neighbors); i++ {
		if g.Nodes[n.Neighbors[i]-1].Marked == false {
			breadthQueue = append(breadthQueue, &g.Nodes[n.Neighbors[i]-1])
		}
	}
}

func Dequeue(n *Node) {
	var tempQueue []*Node
	fmt.Print(n.Num, " ")
	for i := 0; i < len(breadthQueue); i++ {
		if breadthQueue[i] != n {
			tempQueue = append(tempQueue, breadthQueue[i])
		}
	}

	breadthQueue = tempQueue
}

func ClearQueue() {
	breadthQueue = nil
}

func BreadthFirstTraversal(g Graph, n *Node) {
	Visit(n)
	Mark(n)
	Enqueue(&g, n)
	for len(breadthQueue) != 0 {
		Dequeue(n)

		for i := 0; i < len(n.Neighbors); i++ {
			if g.Nodes[n.Neighbors[i]-1].Marked == false {
				Visit(&g.Nodes[n.Neighbors[i]-1])
				Mark(&g.Nodes[n.Neighbors[i]-1])
				Enqueue(&g, &g.Nodes[n.Neighbors[i]-1])
			}
		}
		if len(breadthQueue) != 0 {
			n = breadthQueue[0]
		}
	}
}

func DoDepthTraversal() {
	for i := 1; i <= 4; i++ {
		fmt.Print("\nDepth Traversal Graph ", i, ": ")
		g := GetGraph(i)
		DepthFirstTraversal(g, &g.Nodes[0])
	}
}

func DoBreadthTraversal() {
	for i := 1; i <= 4; i++ {
		fmt.Print("\nBreadth Traversal Graph ", i, ": ")
		g := GetGraph(i)
		BreadthFirstTraversal(g, &g.Nodes[0])
	}
}
