package main

import (
	"fmt"

	"github.com/TOsmanov/my-hw/hw02_fix_app/printer"
	"github.com/TOsmanov/my-hw/hw02_fix_app/reader"
	"github.com/TOsmanov/my-hw/hw02_fix_app/types"
)

func main() {
	var path string

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path)

	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		printer.PrintStaff(staff)
	}
}
