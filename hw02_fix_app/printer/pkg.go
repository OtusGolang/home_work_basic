package printer

import (
	"fmt"

	"github.com/fixme_my_friend/hw02_fix_app/types"
)

// PrintStaff выводит информацию о сотруднике
func PrintStaff(employee types.Employee) {
	// Формируем строку с информацией о сотруднике
	str := fmt.Sprintf("User ID: %d; Age: %d; Name: %s; Department ID: %d",
		employee.UserID, employee.Age, employee.Name, employee.DepartmentID)

	fmt.Println(str)
}
