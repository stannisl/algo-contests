package main

import (
	"bufio"
	"fmt"
	"os"
)

type SageSegmentTree struct {
	n     int
	array []int
	ST    []int
}

func (sst *SageSegmentTree) Init(n int, array []int) *SageSegmentTree {
	sst.n = n
	sst.array = array
	sst.ST = make([]int, 4*n)
	sst.build(1, 0, n-1)

	return sst
}

// func (sst *SageSegmentTree) operation(x []int, y []int) []int {
// 	if x[0] < y[0] {
// 		return x
// 	} else if y[0] < x[0] {
// 		return y
// 	}
// 	return []int{x[0], x[1] + y[1]}
// }

func (sst *SageSegmentTree) operation(x int, y int) int {
	return x + y
}

func (sst *SageSegmentTree) build(node int, L int, R int) {
	if L == R {
		sst.ST[node] = sst.array[L]
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

	var tree SageSegmentTree
	tree.Init(n, array)

	for i := 0; i < 2; i++ {
		var choice, l, r int
		fmt.Fscan(in, &choice, &l, &r)

		if choice == 1 {
			tree.update(1, 0, n-1, l, r)
		}
		if choice == 2 {
			fmt.Fprintln(out, tree.countLogSum(1, 0, n-1, l, r-1))
		}
	}
}

func (sst *SageSegmentTree) update(node int, currentL, currentR, i, v int) {
	if currentR == currentL {
		sst.ST[node] = v
		return
	}
	mid := (currentL + currentR) / 2
	if i <= mid {
		sst.update(2*node, currentL, mid, i, v)
	} else {
		sst.update(2*node+1, mid+1, currentR, i, v)
	}
	sst.ST[node] = sst.ST[2*node] + sst.ST[2*node+1]
}

func (sst *SageSegmentTree) countLogSum(node, currentL, currentR, l, r int) int {
	if currentR < l || currentL > r {
		return 0
	}
	if l <= currentL && currentR <= r {
		return sst.ST[node]
	}
	mid := (currentL + currentR) / 2
	left := sst.countLogSum(2*node, currentL, mid, l, r)
	right := sst.countLogSum(2*node+1, mid+1, currentR, l, r)
	return left + right
}
