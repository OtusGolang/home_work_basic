package printer

import (
	"fmt"

	"github.com/AlexSH61/home_work_basic/types"
)

func PrintStaff(staff []types.Employee) {
	for i := 0; i < len(staff); i++ {
		fmt.Println(staff[i])
	}
}
