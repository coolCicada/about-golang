package main

import (
	"fmt"
	"time"
)

func f(form string) {
	for i := 0; i < 10; i ++ {
		fmt.Println(form, ":", i)
	}
}

func main() {
	f("direct")	

	go f("goroutine")

	go func(msg string) {
		for i := 0; i < 10; i ++ {
			fmt.Println("id:", i, msg)
		}
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
