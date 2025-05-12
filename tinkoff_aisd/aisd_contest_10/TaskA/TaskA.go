package main

import (
	"bufio"
	"fmt"
	"os"
)

type DSU struct {
	ranks  []int
	parent []int

	min  []int
	max  []int
	size []int
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

	ds.size = make([]int, n)
	for i := 0; i < n; i++ {
		ds.size[i] = 1
	}

	ds.max = make([]int, n)
	for i := 0; i < n; i++ {
		ds.max[i] = i
	}

	ds.min = make([]int, n)
	for i := 0; i < n; i++ {
		ds.min[i] = i
	}

	return &ds
}

func (ds *DSU) getMax(x int) int  { return ds.max[ds.find(x)] }
func (ds *DSU) getMin(x int) int  { return ds.min[ds.find(x)] }
func (ds *DSU) getSize(x int) int { return ds.size[ds.find(x)] }

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
	pset := ds.find(x)

	if ds.min[xset] > ds.min[yset] {
		ds.min[pset] = ds.min[yset]
	} else {
		ds.min[pset] = ds.min[xset]
	}

	if ds.max[xset] > ds.max[yset] {
		ds.max[pset] = ds.max[xset]
	} else {
		ds.max[pset] = ds.max[yset]
	}

	ds.size[pset] = ds.size[yset] + ds.size[xset]

}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	dsu := NewDSU(n)

	for i := 0; i < m; i++ {
		var operation string
		var u, v int
		fmt.Fscan(in, &operation)
		if operation == "get" {
			fmt.Fscan(in, &u)
			fmt.Fprintln(out, dsu.getMin(u), dsu.getMax(u), dsu.getSize(u))
		} else if operation == "union" {
			fmt.Fscan(in, &u, &v)
			dsu.union(u, v)
		}
	}
}
