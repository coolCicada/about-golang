package main

import "fmt"

func main() {
	var n int
	var s, t string
	fmt.Scan(&n)
	fmt.Scan(&s, &t)

	for i := 0; i < n; i ++ {
		if s[i] != t[i] {
			if s[i] == '1' && t[i] == 'l' || s[i] == 'l' && t[i] == '1' {
				continue
			} else if s[i] == '0' && t[i] == 'o' || s[i] == 'o' && t[i] == '0' {
				continue
			} else {
				fmt.Println("No")
				return
			}
		}
	}

	fmt.Println("Yes")
}