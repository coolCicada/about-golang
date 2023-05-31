package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 扩容
	var s []int
	for i := 0; i < 3; i ++ {
		s = append(s, i)
	}	
	fmt.Println(s, len(s), cap(s))
	// modifySlice(s)
	// modifySlice2(s)
	modifySlice3(s)
	fmt.Println(s, len(s), cap(s))

	// 初始化
	var j []int // slice 为 nil
	b, _ := json.Marshal(j)
	fmt.Println(string(b))

	j2 := []int{} // slice 不为 nil
	b, _ = json.Marshal(j2)
	fmt.Println(string(b))

	// 优化下标访问
	bce([]int{}) // 编译器优化
}

func normal(s []int) {
	i := 0
	i += s[0]
	i += s[1]
	i += s[2]
	i += s[3]
	println(i)
}

func bce(s []int) {
	_ = s[3]

	i := 0
	i += s[0]
	i += s[1]
	i += s[2]
	i += s[3]
	println(i)
}

func modifySlice(s []int) {
	s = append(s, 2048)
	s[0] = 1024
}


func modifySlice2(s []int) {
	s = append(s, 2048)
	s = append(s, 4096)
	s[0] = 1024
}

func modifySlice3(s []int) {
	s[0] = 1024
	s = append(s, 2048)
	s = append(s, 4096)
}