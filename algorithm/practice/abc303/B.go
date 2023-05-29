package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n ,&m)
	
	r := make([][]int, n + 1)
	for i := 0; i < n + 1; i ++ {
		r[i] = make([]int, n + 1)
	}

	for i := 0; i < m; i ++ {
		a := make([]int, n)
		fmt.Scan(&a[0])
		for j := 1; j < n; j ++ {
			fmt.Scan(&a[j])
			x, y := a[j - 1], a[j]
			r[x][y] = 1
			r[y][x] = 1
		}
	}

	cnt := 0

	for i := 1; i <= n; i ++ {
		for j := 1; j <= n; j ++ {
			if i == j {
				continue
			}
			if r[i][j] == 0 {
				cnt ++
			}
		}
	}

	fmt.Println(cnt / 2)
}