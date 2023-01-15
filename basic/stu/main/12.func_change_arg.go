package main

import "fmt"

func main() {
	/* find(89, 89, 90, 95) */
	// 可变函数传入切片
	/* find(89, []int{89, 90, 91}...) */
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}

func change(s ...string) {
	s = append(s, "playground")
	s[0] = "Go"
	fmt.Println(s)
}

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T \n", nums);
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}

	if !found {
		fmt.Println(num, "not found in", nums)
	}
}