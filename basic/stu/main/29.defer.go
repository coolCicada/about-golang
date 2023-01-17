package main

import "fmt"

func main() {
	// 延迟函数
	/* nums := []int{ 78, 109, 2, 563, 300 }
	largest(nums) */

	// 延迟方法
	/* p := person { "l", "s" }
	defer p.fullName()
	fmt.Printf("Welcome ") */

	// 实参取值
	/* a := 5
	defer printA(a)
	a = 10
	fmt.Println("value of a before deferred function call", a) */

	// defer 栈
	name := "Naveen"
	fmt.Println(name)
	fmt.Println("reverse string:")
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}
}

func printA(a int) {
	fmt.Println("value of a in deferred function", a)
}

type person struct {
	firstName string
	lastName string
}

func (p person) fullName() {
	fmt.Printf("%s %s", p.firstName, p.lastName)
}

func finished() {
	fmt.Println("finished")
}

func largest(nums []int) {  
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
			if v > max {
					max = v
			}
	}
	fmt.Println("Largest number in", nums, "is", max)
}