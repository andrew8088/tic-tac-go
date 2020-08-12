package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("tttapi")
}

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

func PlayRandomMove(seed int64, moves Moves) Moves {
	rand.Seed(seed)

	if !contains(moves, -1) {
		return moves // game is complete
	}

	newMove := rand.Intn(9)

	for contains(moves, newMove) {
		newMove = rand.Intn(9)
	}

	for i, m := range moves {
		if m == -1 {
			moves[i] = newMove
			break
		}
	}

	return moves
}

func contains(moves Moves, item int) bool {
	for _, m := range moves {
		if m == item {
			return true
		}
	}
	return false
}
