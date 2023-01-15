package main

import "fmt"

func main() {
	// break
	/* for i := 1; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Printf(" %d", i);
	} */

	// continue
	/* for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Printf(" %d", i)
	} */

	// 更多例子
	/* i := 0
	for i <= 10 {
		fmt.Printf(" %d ", i)
		i += 2
	} */

	// 无限循环
	i := 0
	for {
		if i >= 100 {
			break;
		}
		i += 2
		fmt.Printf("%d ", i);
	}
}
