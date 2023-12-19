package main

import (
	"fmt"

	"github.com/fixme_my_friend/hw02_fix_app/printer"
	"github.com/fixme_my_friend/hw02_fix_app/reader"
)

func main() {
	var path = "/home/fair/github/data.json"

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	///var err error
	///var staff []types.Employee

	if len(path) == 0 {
		path = "/home/fair/github/data.json"
	}

	var staff, err = reader.ReadJSON(path)
	if err != nil {
		fmt.Print(err)
	}
	///fmt.Print(err)

	printer.PrintStaff(staff)
}
