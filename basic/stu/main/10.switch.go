package main

import "fmt"

func main() {
	switch finger := 5; finger {
	case 1:
		fmt.Println(1)
	case 2, 3, 4, 5:
		fmt.Println("??")
	default:
		fmt.Println("default")
	}

	switch num := number(); {
	case num < 50:
		fmt.Println("< 50")
		fallthrough
	case num < 100:
		fmt.Println("< 100")
		fallthrough
	case num < 200:
		fmt.Println("< 200")
	}
}

func number() int {
	num := 15 * 5
	return num
}
