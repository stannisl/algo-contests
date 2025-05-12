package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type DSU struct {
	ranks  []int
	parent []int
}

type Edge struct {
	u, v int
	cost int
	dir  int
	i, j int
}

func NewDSU(n int) *DSU {
	n++
	var ds = DSU{}

	ds.ranks = make([]int, n)
	for i := 0; i < n; i++ {
		ds.ranks[i] = 1
	}

	ds.parent = make([]int, n)
	for i := 0; i < n; i++ {
		ds.parent[i] = i
	}

	return &ds
}

func (ds *DSU) find(x int) int {
	if ds.parent[x] != x {
		ds.parent[x] = ds.find(ds.parent[x])
	}
	return ds.parent[x]
}

func (ds *DSU) union(x, y int) {
	xnotConnected := ds.find(x)
	ynotConnected := ds.find(y)

	if ynotConnected == xnotConnected {
		return
	}
	if ds.ranks[xnotConnected] < ds.ranks[ynotConnected] {
		ds.parent[xnotConnected] = ynotConnected
	} else if ds.ranks[xnotConnected] > ds.ranks[ynotConnected] {
		ds.parent[ynotConnected] = xnotConnected
	} else {
		ds.parent[ynotConnected] = xnotConnected
		ds.ranks[xnotConnected] = ds.ranks[ynotConnected] + 1
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	board := make([][]int, n)
	dsu := NewDSU((n + 1) * (m + 1))
	for i := 0; i < n; i++ {
		board[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &board[i][j])
		}
	}

	var edges []Edge

	// for i := 0; i < n; i++ {
	// 	for j := 0; j < m; j++ {
	// 		currentNode := i*(m+1) + j
	// 		leftNode := i*(m+1) + (j + 1)
	// 		belowNode := (i+1)*(m+1) + j

	// 		if board[i][j] == 1 || board[i][j] == 3 {
	// 			dsu.union(currentNode, belowNode)
	// 		} else {
	// 			edges = append(edges, Edge{
	// 				u:    currentNode,
	// 				v:    belowNode,
	// 				cost: 1,
	// 				dir:  1,
	// 				i:    i + 1,
	// 				j:    j + 1,
	// 			})
	// 		}

	// 		// Горизонтальная связь (вправо)
	// 		if board[i][j] == 2 || board[i][j] == 3 {
	// 			dsu.union(currentNode, leftNode)
	// 		} else {
	// 			// Добавляем возможное горизонтальное ребро стоимостью 2
	// 			edges = append(edges, Edge{
	// 				u:    currentNode,
	// 				v:    leftNode,
	// 				cost: 2,
	// 				dir:  2,
	// 				i:    i + 1,
	// 				j:    j + 1,
	// 			})
	// 		}
	// 	}
	// }
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			currentNode := i*(m+1) + j

			if i < n-1 {
				belowNode := (i+1)*(m+1) + j
				if board[i][j] == 1 || board[i][j] == 3 {
					dsu.union(currentNode, belowNode)
				} else {
					edges = append(edges, Edge{
						u:    currentNode,
						v:    belowNode,
						cost: 1,
						dir:  1,
						i:    i + 1,
						j:    j + 1,
					})
				}
			}

			if j < m-1 {
				rightNode := i*(m+1) + (j + 1)
				if board[i][j] == 2 || board[i][j] == 3 {
					dsu.union(currentNode, rightNode)
				} else {
					edges = append(edges, Edge{
						u:    currentNode,
						v:    rightNode,
						cost: 2,
						dir:  2,
						i:    i + 1,
						j:    j + 1,
					})
				}
			}
		}
	}
	sort.SliceStable(edges, func(i, j int) bool {
		return edges[i].cost < edges[j].cost
	})

	totalCost := 0
	totalAdded := 0
	var addedEdges []Edge

	for i := 0; i < len(edges); i++ {
		startNode := edges[i].u
		endNode := edges[i].v
		// if edges[i].i == 1 && edges[i].j == 1 {
		// 	fmt.Println(dsu.find(startNode) != dsu.find(endNode))
		// }
		if dsu.find(startNode) != dsu.find(endNode) {
			dsu.union(startNode, endNode)
			totalCost += edges[i].cost
			totalAdded += 1
			addedEdges = append(addedEdges, edges[i])
		}
	}

	fmt.Fprintln(out, totalAdded, totalCost)
	for i := 0; i < len(addedEdges); i++ {
		fmt.Fprintln(out, addedEdges[i].i, addedEdges[i].j, addedEdges[i].dir)
	}
}
