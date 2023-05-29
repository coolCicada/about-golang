package main

import "fmt"

const L int64 = 2e5

type pp struct {
	x, y int64
}

func main() {
	var n, m, h, k int64
	fmt.Scan(&n, &m, &h, &k)

	var s string
	fmt.Scan(&s)

	mp := map[pp]bool{}

	for m > 0 {
		m --
		var x, y int64
		fmt.Scan(&x, &y)
		mp[pp{x, y}] = true
	}

	var ix, iy int64 = 0, 0

	for _, v := range s {
		if h == 0 {
			fmt.Println("No")
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

	fmt.Println("Yes")
}
