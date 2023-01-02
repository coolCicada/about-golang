package main

import "fmt"

func main() {
	find(89, 89, 90, 95)
	find(89, []int{89, 90, 91}...)
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}

func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
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