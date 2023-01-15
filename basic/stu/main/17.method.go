package main

import "fmt"

func main() {
	e := Employee { "ls", 100, "北京", 20 }
	// 方法
	/* e.displaySalary() */

	// 指针接收器与值接收器
	fmt.Println(e.name, e.age)
	e.changeName("newName!!!");
	fmt.Println(e.name, e.age)

	e.changeAge(100)
	fmt.Println(e.age, e.name);
}

type Employee struct {
	name string
	salary int
	currency string
	age int
}

func (e Employee) displaySalary() {
	fmt.Println(e.name, e.currency, e.salary)
}

func (e Employee) changeName(newName string) {
	e.name = newName
}

func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}
