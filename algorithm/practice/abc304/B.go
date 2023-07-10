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

func main() {
	var N string
	fmt.Fscan(r, &N)

	r := ""

	for i := 0; i < len(N); i ++ {
		if i < 3 {
			r += string(N[i])
		} else {
			r += string('0')
		}
	}

	fmt.Fprintln(w, r)

	w.Flush()
}