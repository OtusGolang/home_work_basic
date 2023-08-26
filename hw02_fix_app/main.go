package main

import (
	"fmt"

	"github.com/AlexSH61/home_work_basic/printer"
	"github.com/AlexSH61/home_work_basic/reader"
)

func main() {
	var path string

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	if len(path) == 0 {
		path = "data.json"
	}
	staff, err := reader.ReadJSON(path)
	if err != nil {
		fmt.Println(err)
	} else {
		printer.PrintStaff(staff)
	}
}
