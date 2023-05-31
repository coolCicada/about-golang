package main

import "fmt"

func main() {
	// map 在函数参数传递实际上是指针
	m := make(map[int]int)
	modifyMap(m)
	fmt.Println(m)

	// map 删除key 不会缩容
}

func modifyMap(m map[int]int) {
	m[1] = 1
	m[2] = 2
}