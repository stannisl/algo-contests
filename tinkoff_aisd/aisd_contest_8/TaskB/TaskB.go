package main

import (
	"bufio"
	"fmt"
	"os"
)

type MinSegmentTree struct {
	n     int
	array []int
	ST    [][]int
}

func (sst *MinSegmentTree) Init(n int, array []int) *MinSegmentTree {
	sst.n = n
	sst.array = array
	sst.ST = make([][]int, 4*n)
	sst.build(1, 0, n-1)

	return sst
}

func (sst *MinSegmentTree) operation(x []int, y []int) []int {
	if x[0] > y[0] {
		return y
	} else if y[0] > x[0] {
		return x
	}
	return []int{x[0], x[1] + y[1]}
}

func (sst *MinSegmentTree) build(node int, L int, R int) {
	if L == R {
		sst.ST[node] = []int{sst.array[L], 1}
	} else {
		mid := (L + R) / 2
		sst.build(2*node, L, mid)
		sst.build(2*node+1, mid+1, R)
		sst.ST[node] = sst.operation(sst.ST[2*node], sst.ST[2*node+1])
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	var t int
	fmt.Fscan(in, &n, &t)

	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &array[i])
	}

	var tree MinSegmentTree
	tree.Init(n, array)

	for i := 0; i < t; i++ {
		var choice, l, r int
		fmt.Fscan(in, &choice, &l, &r)

		if choice == 1 {
			tree.update(1, 0, n-1, l, r)
		}
		if choice == 2 {
			res := tree.countMinSum(1, 0, n-1, l, r-1)
			fmt.Fprintln(out, res[0], res[1])
		}
	}
}

func (sst *MinSegmentTree) update(node int, currentL, currentR, i, v int) {
	if currentR == currentL {
		sst.ST[node] = []int{v, 1}
		return
	}
	mid := (currentL + currentR) / 2
	if i <= mid {
		sst.update(2*node, currentL, mid, i, v)
	} else {
		sst.update(2*node+1, mid+1, currentR, i, v)
	}
	sst.ST[node] = sst.operation(sst.ST[2*node], sst.ST[2*node+1])
}

func (sst *MinSegmentTree) countMinSum(node int, currentL, currentR, l, r int) []int {
	if currentR < l || currentL > r {
		return []int{10e10, 10e10}
	}
	if l <= currentL && currentR <= r {
		return sst.ST[node]
	}
	mid := (currentL + currentR) / 2
	left := sst.countMinSum(2*node, currentL, mid, l, r)
	right := sst.countMinSum(2*node+1, mid+1, currentR, l, r)
	return sst.operation(left, right)
}
