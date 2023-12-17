package main

import (
	"fmt"

	"github.com/aasdhajkshd/home_work_basic/tree/master/hw02_fix_app/printer"
	"github.com/aasdhajkshd/home_work_basic/tree/master/hw02_fix_app/reader"
	"github.com/aasdhajkshd/home_work_basic/tree/master/hw02_fix_app/types"
)

func main() {
	path := "data.json"

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	} else {
		fmt.Println("Successfully read data.json")
	}

	staff, err = reader.ReadJSON(path)

	if err != nil {
		fmt.Print(err)
	}

	printer.PrintStaff(staff)
}
