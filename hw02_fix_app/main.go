package main

import (
	"fmt"

	"github.com/fixme_my_friend/hw02_fix_app/printer"
	"github.com/fixme_my_friend/hw02_fix_app/reader"
)

func main() {
	path := "data.json"

	fmt.Printf("Enter data file path (press Enter to use default): ")
	fmt.Scanln(&path)

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err := reader.ReadJSON(path, -1)
	if err != nil {
		fmt.Printf("Error reading JSON: %v\n", err)
		return
	}

	printer.PrintStaff(staff)
}

/*
package main

import (
	"fmt"

	"github.com/fixme_my_friend/hw02_fix_app/printer"
	"github.com/fixme_my_friend/hw02_fix_app/reader"
	"github.com/fixme_my_friend/hw02_fix_app/types"
)

func main() {
	var path = "data.json"

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path, -1)

	fmt.Print(err)

	printer.PrintStaff(staff)
}
*/
