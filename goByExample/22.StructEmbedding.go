package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
	num int
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
		num: 2,
	}	

	fmt.Printf("co = {num: %v, str: %v}\n", co.base.num, co.str)

	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}