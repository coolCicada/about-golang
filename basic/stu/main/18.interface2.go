package main

import "fmt"

func main() {
	// 实现接口：指针接受者与值接受者
	/* var d1 Describer
	p1 := Person{ "SAM", 25 }
	d1 = p1
	p1.Describer()
	p2 := Person{ "james", 32 }
	d1 = &p2
	d1.Describer()

	var d2 Describer
	a := Address{ "wasss", "usa" }
	d2 = &a
	d2.Describer() */

	// 实现多个接口
	
}

type Describer interface {
	Describer()
}

type Person struct {
	name string
	age int
}

func (p Person) Describer() {
	fmt.Println(p.name, p.age)
}

type Address struct {
	state string
	country string
}

func (a *Address) Describer() {
	fmt.Println(a.state, a.country)
}