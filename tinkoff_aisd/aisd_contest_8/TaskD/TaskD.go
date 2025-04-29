package main

import (
	"bufio"
	"fmt"
	"os"
)

type MaxSegmentTree struct {
	n     int
	array []int
	ST    []int
}

func (st *MaxSegmentTree) Init(n int, array []int) *MaxSegmentTree {
	st.n = n
	st.array = array
	st.ST = make([]int, 4*n)
	st.build(1, 0, n-1)

	return st
}

func (st *MaxSegmentTree) operation(x int, y int) int {
	return x + y
}

func (st *MaxSegmentTree) build(node int, L int, R int) {
	if L == R {
		st.ST[node] = st.array[L]
	} else {
		mid := (L + R) / 2
		st.build(2*node, L, mid)
		st.build(2*node+1, mid+1, R)
		st.ST[node] = st.operation(st.ST[2*node], st.ST[2*node+1])
	}
}

func (st *MaxSegmentTree) update(node int, currentL, currentR, i, v int) {
	if currentR == currentL {
		st.ST[node] = v
		return
	}
	mid := (currentL + currentR) / 2
	if i <= mid {
		st.update(2*node, currentL, mid, i, v)
	} else {
		st.update(2*node+1, mid+1, currentR, i, v)
	}
	st.ST[node] = st.ST[2*node] + st.ST[2*node+1]
}

func (st *MaxSegmentTree) query(node, currentL, currentR, l, r, x int) int {
	if currentL > r || currentR < l {
		return 1e18
	}

	if st.ST[node] < x {
		return 1e18
	}

	if currentL == currentR {
		return currentL
	}

	mid := (currentL + currentR) / 2
	left := st.query(2*node, currentL, mid, l, r, x)
	if left != 1e18 {
		return left
	}
	return st.query(2*node+1, mid+1, currentR, l, r, x)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
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

	tree := &MaxSegmentTree{}
	tree.Init(n, array)

	for i := 0; i < t; i++ {
		var choice, l, r int
		fmt.Fscan(in, &choice, &l, &r)

		if choice == 1 {
			tree.update(1, 0, n-1, l, r)
			tree.array[l] = r
		}
		if choice == 2 {
			res := tree.query(1, 0, n-1, r, n-1, l)
			if res != 1e18 {
				fmt.Fprintln(out, res)
			} else {
				fmt.Fprintln(out, -1)
			}
		}
	}
}
