package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	io := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(io, &n)
	array := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(io, &array[i])
	}
	res := 0
	for j := 0; j < n; j++ {
		aj := array[j]

		ai := 0
		for i := 0; i < j; i++ {
			if array[i] > aj {
				ai++
			}
		}

		ak := 0
		for k := j + 1; k < n; k++ {
			if array[k] < aj {
				ak++
			}
		}
		res += ai * ak
		// fmt.Fprintf(out, "j = %d, ai = %d, ak = %d, so res = %d\n", j, ai, ak, res)
	}
	fmt.Fprintln(out, res)

}
