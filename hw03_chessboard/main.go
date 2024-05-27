package main

import (
	"fmt"
	"github.com/AndrewSorokin540/home_work_basic/hw03_chessboard/drawer"
)

func main() {
	var size int
	fmt.Printf("Enter chess board size: ")
	_, scanErr := fmt.Scanf("%d", &size)

	if scanErr != nil {
		fmt.Printf("Error: %v", scanErr)
		return
	}

	chessBoardDrawer.DrawChessBoard(size)
}
