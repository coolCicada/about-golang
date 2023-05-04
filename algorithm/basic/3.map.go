package main

import "fmt"

func main() {
	m := map[string]int{}

	m["hello"] = 1
	m["world"] = 2
	m["hh"] = 3
	fmt.Println(m)

	m["tt"] ++

	delete(m, "hello")

	for k, v := range m {
		println(k, v)
	}
}