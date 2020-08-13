package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("tttapi")
}

func FormatBoard(game Game) string {
	return fmt.Sprintf("%s", game.ToBoard())
}

func FindWinner(game Game) string {
	board := game.ToBoard()
	lines := allLines(board)

	for _, line := range lines {
		if line[0] != "" && enoughMatches(3, line[:]...) {
			return line[0]
		}
	}

	return ""
}

func PlayRandomMove(seed int64, game Game) Game {
	rand.Seed(seed)

	if !contains(game, -1) {
		return game // game is complete
	}

	move := rand.Intn(9)

	for contains(game, move) {
		move = rand.Intn(9)
	}

	for i, m := range game {
		if m == -1 {
			game[i] = move
			break
		}
	}

	return game
}

func FindBlockingMove(game Game) int {
	return 1
}

func contains(game Game, item int) bool {
	for _, m := range game {
		if m == item {
			return true
		}
	}
	return false
}

func enoughMatches(minMatchingCount int, items ...string) bool {
	matchingCount := 0

	for _, item := range items {
		if items[0] == item {
			matchingCount++
		}
	}

	return matchingCount == minMatchingCount
}

func allLines(board Board) [8][3]string {
	lines := [8][3]string{}
	idx := 0

	// rows
	startOfRows := []int{0, 3, 6}
	for _, i := range startOfRows {
		lines[idx] = [3]string{board[i], board[i+1], board[i+2]}
		idx += 1
	}

	// columns
	startOfColumns := []int{0, 1, 2}
	for _, i := range startOfColumns {
		lines[idx] = [3]string{board[i], board[i+3], board[i+6]}
		idx += 1
	}

	// diagonals
	lines[idx] = [3]string{board[0], board[4], board[8]}
	idx += 1
	lines[idx] = [3]string{board[2], board[4], board[6]}

	return lines
}
