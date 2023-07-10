package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
)

type pii struct {
	x, y int
}

var (
	r  = bufio.NewReader(os.Stdin)
	w  = bufio.NewWriter(os.Stdout)
)

func main() {
	var N, D int
	fmt.Fscan(r, &N, &D)

	arr := make([]pii, N)
	vis := make([]bool, N)

	vis[0] = true

	for i := 0; i < N; i ++ {
		fmt.Fscan(r, &arr[i].x, &arr[i].y)
	}

	qu := []pii{}
	qu = append(qu, arr[0])

	for len(qu) > 0 {
		curr := qu[0]
		qu = qu[1:]
		x, y := curr.x, curr.y
		for i := 0; i < N; i ++ {
			nx, ny := arr[i].x, arr[i].y
			t := math.Sqrt(math.Pow(float64(x - nx), 2) + math.Pow(float64(y - ny), 2))
			if t <= float64(D) {
				if !vis[i] {
					vis[i] = true
					qu = append(qu, arr[i])
				}
			}
		}
	}

	for i := 0; i < N; i ++ {
		if vis[i] {
			fmt.Fprintln(w, "Yes")
		} else {
			fmt.Fprintln(w, "No")
		}
	}
	w.Flush()
}