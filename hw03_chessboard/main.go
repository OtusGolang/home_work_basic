package main

import (
	"fmt"
)

func main() {
	var x int
	var y int

	fmt.Printf("Enter lines number: ")
	fmt.Scan(&x)
	fmt.Printf("Enter columns number: ")
	fmt.Scan(&y)

	fmt.Println(x, "*", y)

	for i := 0; i < x; {
		for j := 0; j < y; {
			summ := j + i

			if summ%2 == 0 {
				fmt.Print("   ")
			} else {
				fmt.Print(" # ")
			}
			j++
		}
		i++
		fmt.Print("\n")
	}
}

// test repo
