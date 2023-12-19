package types

import "fmt"

type Employee struct {
	UserID       int    `json:"userId"`
	Age          int    `json:"age"`
	Name         string `json:"name"`
	DepartmentID int    `json:"departmentId"`
	FirstName    string `json:"firstName"`
}

func (e Employee) String() string {
	return fmt.Sprintf("UserId: %d; Age: %d; Name: %s; DepartmentId: %d; FirstName: %s;", e.UserID, e.Age, e.Name, e.DepartmentID, e.FirstName)
}
