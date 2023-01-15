package main

import (
	"fmt"
)

func main() {
	// 创建一个切片
	/* a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1: 4]
	fmt.Println(b)

	c := []int{6, 7, 8}
	fmt.Println(c) */

	// 切片的修改
	/* darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i] ++
	}
	fmt.Println("array after", darr) */
	/* numa := [3]int{ 78, 79, 80}
	nums1 := numa[:]
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa) */

	// 切片的长度和容量
	/* fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
  fruitslice := fruitarray[1:3]
	fmt.Printf("%d - %d", len(fruitslice), cap(fruitslice)); */

	// 追加切片元素
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println(len(cars), cap(cars)) // capacity of cars is 3
	cars2 := append(cars, "Toyota")
	fmt.Println(len(cars2), cap(cars2)) // capacity of cars is doubled to 6
	cars2[1] = "change"
	for _, car := range cars {
		fmt.Printf("%s ", car)
	}
	fmt.Println();
	for _, car := range cars2 {
		fmt.Printf("%s ", car)
	}

	// 切片的函数传递
	/* nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)
	fmt.Println("slice after function call", nos) */
}

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}