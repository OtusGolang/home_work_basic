package printer

import (
	"fmt"

	"github.com/fixme_my_friend/hw02_fix_app/types"
)

func PrintStaff(staff []types.Employee) {
	var str string
	for i := 0; i < len(staff); i++ {
		str = fmt.Sprintf("UserId: %d; Age: %d; Name: %s; DepartmentId: %d; FirstName: %s;",
			staff[i].UserID, staff[i].Age, staff[i].Name, staff[i].DepartmentID, staff[i].FirstName)
		fmt.Println(str)
	}

	///fmt.Println(str)
}
