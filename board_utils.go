package main

import "fmt"

func FormatBoard(moves [9]int) string {
	board := Board{}

	for i, move := range moves {
		// zero-value, no move
		if move == 0 {
			break
		}
		if i%2 == 0 {
			board[move-1] = " x "
		} else {
			board[move-1] = " o "
		}
	}

	return fmt.Sprintf("%s", board)
}
