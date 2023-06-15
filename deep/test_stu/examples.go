package test_stu

import "fmt"

func HelloTom() string {
	return "Jerry"
}

func JudgePassLine(score int16) bool {
	if score >= 60 {
		fmt.Println(">=")
		return true
	}
	fmt.Println("<")
	return false
}

