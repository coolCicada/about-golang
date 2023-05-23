package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "12345"	
	num := s[0] - '0'
	num2 := int(s[0] - '0')
	str := s[0]
	println(s, num)
	fmt.Println(s, num, num2, str)
	t := 1 + '1'
	fmt.Println(t)
	fmt.Printf("%T, %T, %T, %T, %T \n", s, num, num2, str, t)
	fmt.Println(s, num, num2, str, t)
	fmt.Println(strconv.Atoi(s))
	fmt.Println(strconv.Itoa(123))
}