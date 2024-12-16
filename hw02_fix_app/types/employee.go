package types

import "fmt"

// Employee представляет собой сотрудника компании
type Employee struct {
	UserID       int    `json:"user_id"`
	Age          int    `json:"age"`
	Name         string `json:"name"`
	DepartmentID int    `json:"department_id"`
}

// String возвращает строку с описанием сотрудника
func (e Employee) String() string {
	return fmt.Sprintf("User ID: %d; Age: %d; Name: %s; Department ID: %d", e.UserID, e.Age, e.Name, e.DepartmentID)
}
