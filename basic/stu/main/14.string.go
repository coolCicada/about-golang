package main

import (
	"fmt"
)

func main() {
	// 单独获取字符串的每一个字节
	/* name := "Hello World"
	printBytes(name)
	fmt.Println()
	name = "Señor"
	printBytes(name)
	fmt.Println() */

	// rune
	/* name := "Señor"
	printChars(name) */

	// for range
	/* name := "Señor"
	printCharsAndBytes(name) */

	// 字符串的长度
	/* name := "Señor"
	fmt.Println(utf8.RuneCountInString(name)) */

	// 字符串修改
	h := "hello"
	fmt.Println(mutate([]rune(h)))
}

func printBytes (s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}

func printChars(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func printCharsAndBytes(s string) {
	for index, rune := range s {
		fmt.Printf("%c -> %d \n", rune, index);
	}
}

func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}