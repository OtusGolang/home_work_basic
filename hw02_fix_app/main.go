package main

import (
	"fmt"

	"github.com/kidnation-pd/hw02_fix_app/printer"
	"github.com/kidnation-pd/hw02_fix_app/reader"
	"github.com/kidnation-pd/hw02_fix_app/types"
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

	staff, err = reader.ReadJSON(path, -1)
	if err != nil {
		fmt.Print(err)
	}

	printer.PrintStaff(staff)
}
