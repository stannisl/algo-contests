package main

import (
	"bufio"
	"fmt"
	"os"
)

type DSU struct {
	ranks  []int
	parent []int
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
	xset := ds.find(x)
	yset := ds.find(y)

	if yset == xset {
		return
	}
	if ds.ranks[xset] < ds.ranks[yset] {
		ds.parent[xset] = yset
	} else if ds.ranks[xset] > ds.ranks[yset] {
		ds.parent[yset] = xset
	} else {
		ds.parent[yset] = xset
		ds.ranks[xset] = ds.ranks[yset] + 1
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	n++
	m++

	board := make([][]int, n)
	dsu := NewDSU(n * m)
	for i := 0; i < n; i++ {
		board[i] = make([]int, m)
		for j := 0; i < m; j++ {
			fmt.Fscan(in, &board[i][j])
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; i < m; j++ {
			if board[i][j] == 1 {
				dsu.union(i*n+j, (i+1)*n+j)
			} else if board[i][j] == 2 {
				dsu.union(i*n+j, i*n+j+1)
			} else if board[i][j] == 3 {
				dsu.union(i*n+j, i*n+j+1)
				dsu.union(i*n+j, (i+1)*n+j)
			}
		}
	}
	fmt.Fprintln(out, dsu.find(1*n+1) == dsu.find(1*n+3))
}
