package employee

import "fmt"

type employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func New(firstName string, lastName string, totalLeave int, leavesTaken int) employee {
	e := employee { firstName, lastName, totalLeave, leavesTaken }
	return e
}

func (e employee) LeavesRemaining() {
	fmt.Println(e.FirstName, e.LastName, e.TotalLeaves - e.LeavesTaken)
}