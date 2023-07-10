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
	var n int = 0
	mn := int(1e9 + 1)
	fmt.Fscan(r, &n)
	sarr, aarr := make([]string, n), make([]int, n)
	index := 0
	for i := 0; i < n; i ++ {
		fmt.Fscan(r, &sarr[i], &aarr[i])
		if aarr[i] < mn {
			mn = aarr[i]
			index = i
		}
	}

	for i := 0; i < n; i ++ {
		fmt.Fprintln(w, sarr[(index + i) % n])
	}
	w.Flush()
}
