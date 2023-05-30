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

func solve() {
	var n int
	var s string
	fmt.Fscan(r, &n, &s)

	mp := map[string]bool{}
	for i := 1; i < n; i ++ {
		c := s[i - 1: i + 1]
		mp[c] = true
	}

	fmt.Fprintln(w, len(mp))
}

func main() {
	var t int
	fmt.Fscan(r, &t)
	for t > 0 {
		solve()
		t --
	}
	w.Flush()
}

func fmtT() { fmt.Println("over") }