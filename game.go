package main

import "fmt"

type Game [9]int

func MakeGame(partialGame ...int) Game {
	game := Game{-1, -1, -1, -1, -1, -1, -1, -1, -1}

	for i, move := range partialGame {
		game[i] = move
	}

	return game
}

func (g Game) ToBoard() Board {
	board := Board{-1, -1, -1, -1, -1, -1, -1, -1, -1}

	for i, move := range g {
		if move == -1 { // not a move
			break
		}

		board[move] = i
	}
	return board
}

func (g Game) String() string {
	gameStr := "["

	for i, move := range g {
		if move == -1 { // not a move
			break
		}

		if i != 0 { // not the first move
			gameStr += ", "
		}

		gameStr += fmt.Sprintf("%v", move)
	}

	gameStr += "]"

	return gameStr
}

func (g Game) Lines() [8]Line {
	board := g.ToBoard()
	lines := [8]Line{}
	idx := 0

	// rows
	startOfRows := []int{0, 3, 6}
	for _, i := range startOfRows {
		lines[idx] = Line{
			[3]int{i, i + 1, i + 2},
			[3]int{board[i], board[i+1], board[i+2]},
		}
		idx += 1
	}

	// columns
	startOfColumns := []int{0, 1, 2}
	for _, i := range startOfColumns {
		lines[idx] = Line{
			[3]int{i, i + 3, i + 6},
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

func (game Game) Winner() int {
	lines := game.Lines()

	for _, line := range lines {
		if allMovesBySamePlayer(line.indices) {
			return line.indices[0] % 2
		}
	}

	return -1
}

func allMovesBySamePlayer(indices [3]int) bool {
	mod1 := indices[0] % 2
	mod2 := indices[1] % 2
	mod3 := indices[2] % 2
	return mod1 == mod2 && mod2 == mod3 && mod1 != -1
}
