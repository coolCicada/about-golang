package main

import "fmt"

func main() {
	// 数组的声明
	/* var a [3]int
	a[0] = 12
	a[1] = 22
	a[2] = 32
	fmt.Println(a)

	b := [3]int{12, 22, 32}
	fmt.Println(b)

	c := [...]int{12, 3}
	fmt.Println(c) */

	// 数组是值类型
	/* var a = [...]string{"USA", "China", "India", "Germany", "France"}
	b := a
	b[0] = "Singapore"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b) */
	/* num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function", num)
	changeLocal(num)
	fmt.Println("after passing to function", num)
	fmt.Println(len(num)) */

	// 数组的长度
	/* num := [...]float64{ 67.7, 89.8, 21, 78 }
	for i := 0; i < len(num); i++ {
		fmt.Println(num[i])
	}

	for i, v := range num {
		fmt.Println(i, v)
	} */

	// 多维数组
	a := [3][2]string {
		{ "lion", "itger" },
		{ "cat", "dog" },
		{ "pigeon", "peacock" },
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			fmt.Println(a[i][j]);
		}
	}
}

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)
}
