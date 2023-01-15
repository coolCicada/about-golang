package main

import "fmt"

func main() {
	// 创建命名结构体
	/* emp1 := Employee {
		firstName:	"ls",
		lastName:		"sai",
		salary: 		100,
	}

	emp2 := Employee { "a", "b", 2, 100 }

	fmt.Println(emp1, emp2)
	fmt.Println(emp1.firstName, emp1.lastName, emp1.age, emp1.salary) */

	// 提升字段
	var p Person
	p.age = 50
	p.city = "美国"
	fmt.Println(p)
}

type Employee struct {
	firstName, lastName string
	age, salary 				int 
}

type Address struct {
	city, state string
}

type Person struct {
	name string
	age	 int
	Address
}