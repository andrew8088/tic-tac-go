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
	lines := allLines(game)

	for _, line := range lines {
		fmt.Println(line)
		mod1 := line.indices[0] % 2
		mod2 := line.indices[1] % 2
		mod3 := line.indices[2] % 2

		if all(line.indices[:], isNegOne) {
			continue
		}

		if (mod1 == mod2 && mod3 == -1) {
			return line.moves[2]
		}

		if (mod1 == mod3 && mod2 == -1) {
			return line.moves[1]

		}

		if (mod2 == mod3 && mod1 == -1) {
			return line.moves[0]
		}
	}

	return -1
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
			[3]int{i, i+1, i+2},
			[3]int{board[i], board[i+1], board[i+2]},
		}
		idx += 1
	}

	// columns
	startOfColumns := []int{0, 1, 2}
	for _, i := range startOfColumns {
		lines[idx] = Line{
			[3]int{i, i+3, i+6},
			[3]int{board[i], board[i+3], board[i+6]},
		}
		idx += 1
	}

	// diagonals
	lines[idx] = Line{
		[3]int{0, 4, 8},
		[3]int{board[0], board[4], board[8]},
	}
	idx += 1
	lines[idx] = Line{
		[3]int{2, 4, 6},
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

func isNegOne(num int) bool {
	return num == -1
}

type Predicate func (int) bool

func any(arr []int, predicate Predicate) bool {
	for _, m := range arr {
		if predicate(m) {
			return true
		}
	}
	return false
}

func all(arr []int, predicate Predicate) bool {
	for _, m := range arr {
		if !predicate(m) {
			return false
		}
	}
	return true
}
