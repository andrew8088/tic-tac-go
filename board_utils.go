package main

import "fmt"

func FormatBoard(moves Moves) string {
	return fmt.Sprintf("%s", moves.ToBoard())
}

func FindWinner(moves Moves) string {
	board := moves.ToBoard()

	// Check rows
	startOfRows := []int{0, 3, 6}
	for _, i := range startOfRows {
		if board[i] != "" && board[i] == board[i+1] && board[i+1] == board[i+2] {
			return board[i]
		}
	}

	// Check columns
	startOfColumns := []int{0, 1, 2}
	for _, i := range startOfColumns {
		if board[i] != "" && board[i] == board[i+3] && board[i+3] == board[i+6] {
			return board[i]
		}
	}

	// Check diagonals
	if board[0] != "" && board[0] == board[4] && board[4] == board[8] {
		return board[0]
	}

	if board[2] != "" && board[2] == board[4] && board[4] == board[6] {
		return board[2]
	}

	return ""
}
