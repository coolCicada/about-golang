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
	var a = [...]string{"USA", "China", "India", "Germany", "France"}
	b := a
	b[0] = "Singapore"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)

	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function", num)
	changeLocal(num)
	fmt.Println("after passing to function", num)
	fmt.Println(len(num))

	for i := 0; i < len(num); i++ {
		fmt.Println(num[i])
	}

	for i, v := range num {
		fmt.Println(i, v)
	}
}

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)
}
