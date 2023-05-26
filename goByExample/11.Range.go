package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for i, num := range nums {
		fmt.Println(i, num)
		sum += num
	}
	fmt.Println("sum:", sum)

	kvs := map[string]string{"a": "applw", "b": "banana"}
	for k, v := range kvs {
		// fmt.Printf("%s -> %s\n", k, v)
		fmt.Println(k, "->", v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go哈哈" {
		fmt.Println(i, string(c))
	}
}