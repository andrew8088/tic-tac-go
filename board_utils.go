package main

import "fmt"

func FormatBoard(moves Moves) string {
	return fmt.Sprintf("%s", moves.ToBoard())
}

func FindWinner(moves Moves) string {
	board := moves.ToBoard()

	// Check rows
	if board[0] != "" && board[0] == board[1] && board[1] == board[2] {
		return board[0]
	}

	if board[3] != "" && board[3] == board[4] && board[4] == board[5] {
		return board[3]
	}

	if board[6] != "" && board[6] == board[7] && board[7] == board[8] {
		return board[6]
	}

	// Check columns
	if board[0] != "" && board[0] == board[3] && board[3] == board[6] {
		return board[0]
	}

	if board[1] != "" && board[1] == board[4] && board[4] == board[7] {
		return board[1]
	}

	if board[2] != "" && board[2] == board[5] && board[5] == board[8] {
		return board[2]
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
