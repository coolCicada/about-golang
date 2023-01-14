package main

import "fmt"

func main() {
	// 声明单个变量
	/* var age int
	fmt.Println("my age is", age)
	age = 29
	fmt.Println("my age is", age)
	age = 54
	fmt.Println("my new age is", age) */

	// 声明变量并初始化
	/* var age int = 29
	fmt.Println("my age is", age) */

	// 类型推断
	/* var age = 29
	fmt.Println("my age is", age); */

	// 声明多个变量
	/* var width, height int = 100, 50
	fmt.Println("width is", width, "height is", height) */
	/* var (
		name = "ls"
		age = 29
		height int
	)
	fmt.Println("name:", name, "age:", age, "height:", height) */

	// 简短声明 ( := 左边至少有一个未声明的变量 )
	/* a, b := 20, 30
	fmt.Println("a is", a, "b is", b)
	b, c := 40, 50
	fmt.Println("b is", b, "c is", c)
	b, c = 100, 200
	fmt.Println("change b is", b, "c is", c) */
}