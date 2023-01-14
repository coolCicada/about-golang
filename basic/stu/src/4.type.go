package main

import (
	"fmt"
)

func main() {
	// bool
	/* a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d) */

	// 有符号整型
	/* a := 89
	var b int8 = 95
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))
	fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b)) */

	// 浮点型
	/* a, b := 0.1, 0.2
	fmt.Printf("type of a,b %T, %T", a, b)
	fmt.Println("total: ", a+b)
	sum := a + b
	diff := a - b
	fmt.Println("sum", sum, "diff", diff) */

	// string 类型
	/* first := "ls"
	last := "rigt"
	name := first + " " + last
	fmt.Println("My name is", name)
	a := 'c'
	b := 'd'
	fmt.Println(a - b) */

	// 类型转换
	/* i := 58
	j := 67.8
	sum := i + int(j)
	fmt.Println(sum)
	j = float64(i)
	fmt.Printf("%f %T", j, j) */
}
