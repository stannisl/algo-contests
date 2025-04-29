/*
https://coderun.yandex.ru/selections/2024-summer-backend/problems/balls-and-baskets/description?compiler=go
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

type SegmentTree struct {
	n     int
	array []int
	ST    []int
}

const mod = 1e9 + 7

func (st *SegmentTree) init(n int, array []int) {
	st.n = n
	st.array = array
	st.ST = make([]int, 4*n)
	st.build(1, 0, n-1)
}

func (st *SegmentTree) operation(x, y int) int {
	return (x * y) % mod
}

func (st *SegmentTree) build(node, l, r int) {
	if l == r {
		st.ST[node] = st.array[r]
	} else {
		mid := (l + r) / 2
		st.build(2*node, l, mid)
		st.build(2*node+1, mid+1, r)
		st.ST[node] = st.operation(st.ST[2*node], st.ST[2*node+1])
	}
}

func (st *SegmentTree) query(node, currL, currR, l, r int) int {
	if currR < l || currL > r {
		return 1
	}
	if l <= currL && currR <= r {
		return st.ST[node]
	}
	mid := (currL + currR) / 2
	leftres := st.query(2*node, currL, mid, l, r)
	rightres := st.query(2*node+1, mid+1, currR, l, r)
	return (leftres * rightres) % mod
}

func (st *SegmentTree) update(node, currL, currR, l, r int) {

	if currL > r || currR < l {
		return
	}
	if currL == currR {
		st.ST[node] += 1
		st.ST[node] %= mod
		return
	}

	mid := (currR + currL) / 2
	st.update(2*node, currL, mid, l, r)
	st.update(2*node+1, mid+1, currR, l, r)
	st.ST[node] = st.operation(st.ST[2*node], st.ST[2*node+1])
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n)
	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &array[i])
	}

	var st SegmentTree
	st.init(n, array)

	fmt.Fscan(in, &k)

	for i := 0; i < k; i++ {
		var choice, l, r int
		fmt.Fscan(in, &choice, &l, &r)
		if choice == 1 {
			fmt.Fprintln(out, (st.query(1, 0, n-1, l-1, r-1) % mod))
		} else if choice == 0 {
			st.update(1, 0, n-1, l-1, r-1)
		}
	}
}
