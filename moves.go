package main

import "fmt"

type Moves [9]int

func MakeMoves(partialMoves ...int) Moves {
	moves := Moves{-1, -1, -1, -1, -1, -1, -1, -1, -1}

	for i, m := range partialMoves {
		moves[i] = m
	}

	return moves
}

func (m Moves) ToBoard() Board {
	board := Board{}

	for i, move := range m {
		if move == -1 { // not a move
			break
		}
		if i%2 == 0 {
			board[move] = "x"
		} else {
			board[move] = "o"
		}
	}
	return board
}

func (m Moves) String() string {
	moveStr := "["

	for i, move := range m {
		if move == -1 { // not a move
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
