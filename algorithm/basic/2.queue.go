package main

import (
	"fmt"
)

func main() {
	queue := []int{}
	queue = append(queue, 10)
	fmt.Println(queue)

	v := queue[0]
	fmt.Println(v)
	queue = queue[1:]

	fmt.Println(len(queue))	
}