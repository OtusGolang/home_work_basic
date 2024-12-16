package main

import "fmt"

func main() {
	size := 8 // Размер доски (N x N)
	board := createChessBoard(size)
	fmt.Println(board)
}

func createChessBoard(size int) string {
	var board string
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				board += "#"
			} else {
				board += " "
			}
		}
		board += "\n"
	}
	return board
}
