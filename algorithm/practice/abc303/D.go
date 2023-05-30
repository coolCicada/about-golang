package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	r  = bufio.NewReader(os.Stdin)
	w  = bufio.NewWriter(os.Stdout)
)

func main() {
	var x, y, z int64
	fmt.Fscan(r, &x, &y, &z)

	var s string
	fmt.Fscan(r, &s)

	a := make([][]int64, len(s) + 1)
	for i := 0; i <= len(s); i ++ {
		a[i] = make([]int64, 2)
	}
	
	const inf = 1e9

	a[0][1] = inf

	for i := 1; i <= len(s); i ++ {
		if s[i - 1] == 'a' {
			a[i][0] = min(a[i - 1][0] + x, a[i - 1][1] + x + z)
			a[i][1] = min(a[i - 1][1] + y, a[i - 1][0] + y + z)
		} else {
			a[i][0] = min(a[i - 1][0] + y, a[i - 1][1] + y + z)
			a[i][1] = min(a[i - 1][1] + x, a[i - 1][0] + x + z)
		}
	}
	fmt.Fprintln(w, min(a[len(s)][0], a[len(s)][1]))
	w.Flush()
}

func min(a ...int64) int64 {
	mn := a[0]
	for i := 1; i < len(a); i ++ {
		if a[i] < mn {
			mn = a[i]
		}
	}
	return mn
}