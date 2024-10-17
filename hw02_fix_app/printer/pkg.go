package printer

import (
	"fmt"
	"hw02_fix_app/types"
)

func PrintStaff(staff []types.Employee) {
	for i := 0; i < len(staff); i++ {
		str := fmt.Sprintf(
			"User ID: %d; Age: %d; Name: %s; Department ID: %d; ",
			staff[i].UserID, staff[i].Age, staff[i].Name, staff[i].DepartmentID)
		fmt.Println(str)
	}
}
