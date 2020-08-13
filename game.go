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
	board := Board{}

	for i, move := range g {
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
