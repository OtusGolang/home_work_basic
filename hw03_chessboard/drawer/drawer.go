package chessBoardDrawer

import "fmt"

func DrawChessBoard(size int) {
	var boardString string

	for i := 1; i <= size; i++ {
		if i%2 == 0 {
			for j := 1; j <= size; j++ {
				if j%2 == 0 {
					boardString = boardString + string('#')
				} else {
					boardString = boardString + string('_')
				}
			}
		} else {
			for j := 1; j <= size; j++ {
				if j%2 == 0 {
					boardString = boardString + string('_')
				} else {
					boardString = boardString + string('#')
				}
			}
		}

		boardString = boardString + string('\n')
	}

	fmt.Printf(boardString)
}
