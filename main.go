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
	lines := allLines(game)

	for _, line := range lines {
		if allMovesBySamePlayer(line.indices) {
			if line.indices[0]%2 == 0 {
				return "x"
			} else {
				return "o"
			}
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

func allLines(game Game) [8]Line {
	board := game.ToBoard()
	lines := [8]Line{}
	idx := 0

	// rows
	startOfRows := []int{0, 3, 6}
	for _, i := range startOfRows {
		lines[idx] = Line{
			[3]int{board[i], board[i+1], board[i+2]},
		}
		idx += 1
	}

	// columns
	startOfColumns := []int{0, 1, 2}
	for _, i := range startOfColumns {
		lines[idx] = Line{
			[3]int{board[i], board[i+3], board[i+6]},
		}
		idx += 1
	}

	// diagonals
	lines[idx] = Line{
		[3]int{board[0], board[4], board[8]},
	}
	idx += 1
	lines[idx] = Line{
		[3]int{board[2], board[4], board[6]},
	}

	return lines
}

func allMovesBySamePlayer(indices [3]int) bool {
	mod1 := indices[0] % 2
	mod2 := indices[1] % 2
	mod3 := indices[2] % 2
	return mod1 == mod2 && mod2 == mod3 && mod1 != -1
}
