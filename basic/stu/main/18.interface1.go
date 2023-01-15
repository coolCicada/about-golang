package main

import "fmt"

func main() {
	/* name := MyString("Sam Anderson")
	res := name.FindVowels()
	fmt.Printf("%c", res) */

	// 接口的实际用途
	/* pemp1 := Permanent{ 1, 5000, 20 }
	pemp2 := Permanent{ 2, 6000, 30 }
	cemp1 := Contract{ 3, 3000 }
	employees := []SalaryCalculator{ pemp1, pemp2, cemp1 }
	totalExpense(employees) */

	// 类型选择
	findType("Naveen")
	findType(77)
	findType(89.80)
}

func findType(i interface{}) {
	switch i.(type) {
	case string: 
		fmt.Println(i.(string))
	case int:
		fmt.Println(i.(int))
	default:
		fmt.Println("unknown type")
	}
}

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId int
	basicpay int
	pf int
}

type Contract struct {
	empId int
	basicpay int
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("$%d", expense)
}

type VowelsFInder interface {
	FindVowels() []rune
}

type MyString string

func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if (rune == 'a' ||
				rune == 'e' || 
				rune == 'i' ||
				rune == 'o' ||
			  rune == 'u') {
					vowels = append(vowels, rune)
				}
	}
	return vowels
}