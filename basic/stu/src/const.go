package main

import "fmt"

func main() {
	// 定义
	/* const a int = 50
	var b string = "i love Go"
	fmt.Println(a, b)
	fmt.Println(a, b) */

	// 字符串常量
	/* var name = "Sam"
	name = "b"
	fmt.Printf("type %T value %v", name, name) */

	// 数字常量
	const a = 5
	var f float64 = a
	fmt.Printf("type %T \n", a)
	fmt.Println("value:", f, "---")
}
