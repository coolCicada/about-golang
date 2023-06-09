package main

import (
	"fmt"
	"sync"
	"time"
)

func hello(i int) {
	fmt.Println(i)
}

func Test1() {
	for i := 0; i < 5; i ++ {
		go hello(i)	
	}
	time.Sleep(time.Second)
}

func Test2() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i ++ {
		go func(j int) {
			defer wg.Done()
			hello(j)
		}(i)
	}
	wg.Wait()
}

func CalSquare() {
	src := make(chan int)
	dest := make(chan int, 3)
	go func() {
		defer close(src)
		for i := 0; i < 10; i ++ {
			src <- i
		}
	}()
	
	go func() {
		defer close(dest)
		for i := range src {
			time.Sleep(time.Second)
			dest <- i * i
		}
	}()

	for i := range dest {
		println(i)
	}

	println("over")
}
func main() {
	// Test1()
	Test2()
	// CalSquare()
}