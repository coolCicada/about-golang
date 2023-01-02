package main

import (
	"fmt"
)

func main() {
	fmt.Println(calculateBill(10, 2))
	area, _ := rectProps(2.1, 1.2)
	fmt.Println(area)
}

// 示例函数
func calculateBill(price int, no int) int {
	totalPrice := price * no
	return totalPrice
}

// 多返回值
func rectProps(length, width float64) (area, primeter float64) {
	area = length * length
	primeter = (length + width) * 2
	return
}
