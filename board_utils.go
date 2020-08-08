package main

import "fmt"

func FormatBoard(moves Moves) string {
	return fmt.Sprintf("%s", moves.ToBoard())
}

func FindWinner(moves Moves) string {
	// Check rows
	board := moves.ToBoard()

	if board[0] != "" && board[0] == board[1] && board[1] == board[2] {
		return board[0]
	}

	if board[3] != "" && board[3] == board[4] && board[4] == board[5] {
		return board[3]
	}

	if board[6] != "" && board[6] == board[7] && board[7] == board[8] {
		return board[6]
	}

	return ""
}
