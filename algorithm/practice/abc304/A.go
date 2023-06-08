package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	r  = bufio.NewReader(os.Stdin)
	w  = bufio.NewWriter(os.Stdout)
)

type pii struct {
	name string
	age int
}
func main() {
	var n int
	fmt.Fscan(r, &n)
	a := make([]pii, n)
	for i := 0; i < n; i ++ {
		fmt.Fscan(r, &a[i].name, &a[i].age)
	}
	mn := 0
	mage := a[0].age
	for i := 0; i < n; i ++ {
		if a[i].age < mage {
			mage = a[i].age
			mn = i
		}
	}
	for i := 0; i < n; i ++ {
		fmt.Fprintln(w, a[(i + mn) % n].name)
	}
	w.Flush()
}

func fmtT() { fmt.Println("over") }