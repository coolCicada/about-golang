package main

import "mtest/employee"

func main() {
	// e := employee.Employee {
	// 	FirstName: "Sam",
	// 	LastName: "Adolf",
	// 	TotalLeaves: 30,
	// 	LeavesTaken: 20,
	// }
	e := employee.New("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()
}
