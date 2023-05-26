package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
		fmt.Println(s)

		const n = 500000000

		const d = 3e20 / n
		fmt.Println(d)

		fmt.Println(int64(d))

		fmt.Println(math.Sin(n))

		// var n2 int = 10000
		// fmt.Println(math.Sin(n2))  XX
}