package main

import "fmt"

func main() {
	/* b := 25
	a := &b
	fmt.Printf("%T %T \n", a, b)
	fmt.Println(a) */

	// 指针的解引用
	/* b := 255
	a := &b 
	*a ++
	fmt.Println(a, *a, b) */

	// 向函数传递指针参数
	/* a := 58
	fmt.Println(a)
	b := &a 
	change(b)
	fmt.Println(a) */

	// 使用切片而不是指针
	a := [3]int{ 89, 90, 91}
	modify(a[:])
	fmt.Println(a)
}

func change(val *int) {
	*val = 55
}

func modify(sls []int) {
	sls[0] = 90
}