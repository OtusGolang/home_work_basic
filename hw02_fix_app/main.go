package main

import (
	"fmt"

	"github.com/krupovroman/home_work_basic/hw02_fix_app/printer"
	"github.com/krupovroman/home_work_basic/hw02_fix_app/reader"
	"github.com/krupovroman/home_work_basic/hw02_fix_app/types"
)

func main() {
	var path string = "data.json"

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	} else {
	}

	staff, err = reader.ReadJSON(path, -1)

	fmt.Print(err)

	printer.PrintStaff(staff)
}
