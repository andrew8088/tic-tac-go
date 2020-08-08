package main

import "fmt"

type Moves [9]int

func (m Moves) ToBoard() Board {
	board := Board{}

	for i, move := range m {
		// zero-value, no move
		if move == 0 {
			break
		}
		if i%2 == 0 {
			board[move-1] = "x"
		} else {
			board[move-1] = "o"
		}
	}
	return board
}

func (m Moves) String() string {
	moveStr := "["

	for i, move := range m {
		if move == 0 { // not a move
			break
		}

		if i != 0 { // not the first move
			moveStr += ", "
		}

		moveStr += fmt.Sprintf("%v", move)
	}

	moveStr += "]"

	return moveStr
}
