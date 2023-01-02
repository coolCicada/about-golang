package main

import "fmt"

func main() {
	// 例子
	i := 1
	for i <= 10 {
		// if i > 5 {
		// 	break
		// }
		i++

		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
