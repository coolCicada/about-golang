package main

import (
	"fmt"
)

func main() {
	/* a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1: 4]
	fmt.Println(b)

	c := []int{6, 7, 8}
	fmt.Println(c)

	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)

	for i := range dslice {
		dslice[i] ++
	}
	fmt.Println("array after", darr) */

	// 切片的长度和容量
	/* fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
  fruitslice := fruitarray[1:3]
	fmt.Printf("%d - %d", len(fruitslice), cap(fruitslice)); */

	// 追加切片元素
	/* cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) // capacity of cars is 3
	cars2 := append(cars, "Toyota")
	fmt.Println("cars2:", cars2, "has new length", len(cars2), "and capacity", cap(cars2)) // capacity of cars is doubled to 6
	cars2[1] = "change"
	fmt.Println(cars, cars2) */

	// 切片的函数传递
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)
	fmt.Println("slice after function call", nos)
}

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}