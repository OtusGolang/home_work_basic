package main

import (
	"fmt"

	"github.com/fixme_my_friend/hw02_fix_app/printer"
	"github.com/fixme_my_friend/hw02_fix_app/reader"
	"github.com/fixme_my_friend/hw02_fix_app/types"
)

func main() {
	var path string

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	var staff []types.Employee
	var err error

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path, -1)
	if err != nil {
		fmt.Printf("Error reading JSON: %v\n", err)
		return
	}

	// Печать информации о каждом сотруднике
	for _, employee := range staff {
		printer.PrintStaff(employee)
	}
}
