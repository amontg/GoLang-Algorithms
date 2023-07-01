package graphing

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Tree struct {
	Branches []Fringe // permanent nodes that are part of the tree (solid lines)
}

type Fringe struct {
	Main    Node
	ToStart string
	Adj     []Node // list of adjacent nodes
}

type Graph struct {
	Nodes []Node
}

type Node struct {
	Name   string
	Marked bool
	Adj    []string
}

func GetGraph(num int) *Graph {
	var chosenGraph = new(Graph)

	graph1a := Graph{
		Nodes: []Node{
			{Name: "A", Marked: false, Adj: []string{"B,2", "F,3"}},
			{Name: "B", Marked: false, Adj: []string{"A,2", "C,5", "E,4", "G,3"}},
			{Name: "C", Marked: false, Adj: []string{"B,5", "E,2", "H,4"}},
			{Name: "D", Marked: false, Adj: []string{"E,2", "G,2", "F,4"}},
			{Name: "E", Marked: false, Adj: []string{"B,4", "C,2", "H,3", "D,2"}},
			{Name: "F", Marked: false, Adj: []string{"A,3", "D,4", "G,1"}},
			{Name: "G", Marked: false, Adj: []string{"F,1", "H,1", "D,2", "B,3"}},
			{Name: "H", Marked: false, Adj: []string{"C,4", "E,3", "G,1"}},
		},
	}

	switch num {
	case 1:
		chosenGraph = &graph1a
	}

	return chosenGraph
}

func DoDijPrim() {
	var graph = new(Graph)
	graph = GetGraph(1)
	var tree = new(Tree)

	DijPrim(tree, graph, &graph.Nodes[0])
	//PrintTree(tree)
}

func DijPrim(t *Tree, g *Graph, n *Node) {
	CreateBranch(t, g, n)
	var visited []string
	var queue []string
	visited = append(visited, n.Name)

	for i := 0; i < len(n.Adj); i++ {
		for j := 0; j < len(visited); j++ {
			temp := strings.Split(n.Adj[i], ",")
			if strings.Compare(visited[j], temp[0]) != 0 {
				queue = append(queue, n.Adj[i])
			}
		}
	}
	SortQueue(queue)

	var smallest *Node
	for len(visited) < len(g.Nodes) {
		smallest, queue = GetSmallest(t, g, queue)
		to := smallest.Name

		if !VisitedContains(visited, to) {
			CreateBranch(t, g, smallest)
			visited = append(visited, to)

			for i := 0; i < len(smallest.Adj); i++ {
				temp := strings.Split(smallest.Adj[i], ",")
				if !VisitedContains(visited, temp[0]) {
					queue = append(queue, smallest.Adj[i])
				}
			}
		}

		PrintTree(t)
	}
}

func FindNode(g *Graph, name string) *Node {
	for i := 0; i < len(g.Nodes); i++ {
		if strings.Compare(g.Nodes[i].Name, name) == 0 {
			return &g.Nodes[i]
		}
	}

	return nil
}

func VisitedContains(visited []string, target string) bool {
	for i := 0; i < len(visited); i++ {
		if strings.Compare(visited[i], target) == 0 {
			return true
		}
	}

	return false
}

func SortQueue(queue []string) {
	for i := 0; i < len(queue); i++ {
		temp := strings.Split(queue[i], ",")
		tempInt, _ := strconv.Atoi(temp[1])
		temp2 := strings.Split(queue[0], ",")
		temp2Int, _ := strconv.Atoi(temp2[1])
		if tempInt < temp2Int {
			insert(queue, 0, queue[i])
		}
	}
}

func insert(a []string, index int, value string) {
	if len(a) == index { // nil or empty slice or after last element
		a = append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
}

func remove(a []string, target string) []string {
	var tempArray []string
	for i := 0; i < len(a); i++ {
		temp := strings.Split(a[i], ",")
		if strings.Compare(temp[1], target) != 0 {
			tempArray = append(tempArray, a[i])
		}
	}

	return tempArray
}

func GetSmallest(t *Tree, g *Graph, queue []string) (*Node, []string) {
	var snode string
	snum := 100
	for i := 0; i < len(queue); i++ {
		temp := strings.Split(queue[i], ",")
		tempInt, _ := strconv.Atoi(temp[1])
		if tempInt < snum {
			snode = queue[i]
			snum = tempInt
		}
	}

	queue = remove(queue, snode)

	return FindNode(g, snode), queue
}

func GetWeight(from Node, to string) int64 {
	var weight int64
	var err error

	for i := 0; i < len(from.Adj); i++ {
		temp := strings.Split(from.Adj[i], ",")

		if strings.Compare(temp[0], to) == 0 {
			weight, err = strconv.ParseInt(temp[1], 0, 64)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return weight
}

func MoreEdges(g *Graph) bool { // true if there's more nodes to check, false if all have been marked
	check := false

	for i := 0; i < len(g.Nodes); i++ {
		if g.Nodes[i].Marked == false {
			check = true
			break
		}
	}

	return check
}

func CreateBranch(t *Tree, g *Graph, n *Node) {
	var newFringe Fringe
	n.Marked = true
	newFringe.Main = *n
	//newFringe.ToStart = edge
	for i := 0; i < len(n.Adj); i++ {
		temp := strings.Split(n.Adj[i], ",")
		for j := 0; j < len(g.Nodes); j++ {
			if strings.Compare(g.Nodes[j].Name, temp[0]) == 0 {
				if g.Nodes[j].Marked == false {
					newFringe.Adj = append(newFringe.Adj, g.Nodes[j])
				}
				break
			}
		}
	}

	PrintFringe(&newFringe)
	t.Branches = append(t.Branches, newFringe)
}

// EVERYTHING BELOW HERE

func IsMarked(g *Graph, name string) bool {
	for i := 0; i < len(g.Nodes); i++ {
		if strings.Compare(g.Nodes[i].Name, name) == 0 {
			if g.Nodes[i].Marked == true {
				return true
			}
		}
	}

	return false
}

func RemoveSmallest(t *Tree, n *Node, old *Node) {
	var tempAdjSlice []Node
	for i := 0; i < len(t.Branches); i++ {
		if strings.Compare(t.Branches[i].Main.Name, old.Name) == 0 {
			for j := 0; j < len(t.Branches[i].Adj); j++ {
				if strings.Compare(t.Branches[i].Adj[j].Name, n.Name) != 0 {
					tempAdjSlice = append(tempAdjSlice, t.Branches[i].Adj[j])
				}
			}
		}

		t.Branches[i].Adj = tempAdjSlice
	}
}

func DijMin(t *Tree, g *Graph, start, end *Node) {
	if len(t.Branches) == 0 {
		CreateBranch(t, g, start)
	}

	for strings.Compare(start.Name, end.Name) != 0 {
		//oldnode := start
		// start = GetSmallest(t, g,)

		// CreateBranch(t, g, start)

		// DijMin(t, g, start, end)
	}
}

func DoDijMin() {
	var mingraph = new(Graph)
	mingraph = GetGraph(1)
	var mintree = new(Tree)

	DijMin(mintree, mingraph, &mingraph.Nodes[0], &mingraph.Nodes[len(mingraph.Nodes)-1])
}

func PrintTree(t *Tree) {
	for i := 0; i < len(t.Branches); i++ {
		fmt.Println(t.Branches[i])
	}
}

func PrintFringe(fringe *Fringe) {
	fmt.Printf("\n%v -> ", fringe.Main.Name)
	for i := 0; i < len(fringe.Adj); i++ {
		fmt.Printf("(%v)", fringe.Adj[i].Name)
	}
}
