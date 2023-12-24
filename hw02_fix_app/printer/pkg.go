package printer

import (
	"fmt"

	"github.com/slimko/home_work_basic/hw02_fix_app/types"
)

func PrintStaff(staff []types.Employee) {
	var str string
	for i := 0; i < len(staff); i++ {
		str = fmt.Sprintf(
			"User ID: %d; Age: %d; Name: %s; Department ID: %d;",
			staff[i].UserID, staff[i].Age, staff[i].Name, staff[i].DepartmentID,
		)
		fmt.Println(str)
	}
}
