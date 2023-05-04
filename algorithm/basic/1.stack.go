package main

import "fmt"

func main() {
	stack := []int{}
	stack = append(stack, 10)
	fmt.Println(stack)
	
	v := stack[len(stack) - 1]
	fmt.Println(v)
	stack = stack[:len(stack) - 1]

	fmt.Println(len(stack))
}