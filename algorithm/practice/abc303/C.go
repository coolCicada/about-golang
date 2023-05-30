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

const L int64 = 2e5

type pp struct {
	x, y int64
}

func main() {
	solve()
	w.Flush()
}

func solve() {
	var n, m, h, k int64
	fmt.Fscan(r, &n, &m, &h, &k)

	var s string
	fmt.Fscan(r, &s)

	mp := map[pp]bool{}

	for m > 0 {
		m --
		var x, y int64
		fmt.Fscan(r, &x, &y)
		mp[pp{x, y}] = true
	}

	var ix, iy int64 = 0, 0

	for _, v := range s {
		if h == 0 {
			fmt.Fprintln(w, "No")
			return
		}

		h --

		switch(v) {
		case 'R':
			ix ++
		case 'L':
			ix --
		case 'U':
			iy ++
		case 'D':
			iy --
		}

		if h < k && mp[pp{ix, iy}] {
			h = k
			mp[pp{ix, iy}]= false
		}
	}

	fmt.Fprintln(w, "Yes")
}
