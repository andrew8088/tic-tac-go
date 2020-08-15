package main

import "testing"

func TestCanMakeGame(t *testing.T) {
	cases := []struct {
		actual   Game
		expected Game
	}{
		{MakeGame(), Game{-1, -1, -1, -1, -1, -1, -1, -1, -1}},
		{MakeGame(0), Game{0, -1, -1, -1, -1, -1, -1, -1, -1}},
		{MakeGame(0, 8, 1, 7), Game{0, 8, 1, 7, -1, -1, -1, -1, -1}},
	}

	for _, c := range cases {
		Equal(t, c.actual, c.expected)
	}
}

func TestCanConvertGameToBoard(t *testing.T) {
	cases := []struct {
		input    Game
		expected Board
	}{
		{MakeGame(), Board{-1, -1, -1, -1, -1, -1, -1, -1, -1}},
		{MakeGame(0), Board{0, -1, -1, -1, -1, -1, -1, -1, -1}},
		{MakeGame(0, 7), Board{0, -1, -1, -1, -1, -1, -1, 1, -1}},
		{MakeGame(0, 7, 1, 2), Board{0, 2, 3, -1, -1, -1, -1, 1, -1}},
	}

	for _, c := range cases {
		actual := c.input.ToBoard()
		Equal(t, actual, c.expected)
	}
}

func TestCanConvertGameToString(t *testing.T) {
	cases := []struct {
		input    Game
		expected string
	}{
		{MakeGame(), "[]"},
		{MakeGame(1), "[1]"},
		{MakeGame(1, 8), "[1, 8]"},
		{MakeGame(1, 8, 2, 3), "[1, 8, 2, 3]"},
	}

	for _, c := range cases {
		actual := c.input.String()
		Equal(t, actual, c.expected)
	}
}

func TestCanDetermineTheWinner(t *testing.T) {
	const X = 0
	const O = 1

	cases := []struct {
		input    Game
		expected int
	}{
		{MakeGame(), -1},
		// row wins
		{MakeGame(0, 3, 1, 4, 2), X},
		{MakeGame(0, 3, 1, 4, 6, 5), O},
		{MakeGame(8, 3, 7, 4, 6), X},
		// columns wins
		{MakeGame(0, 1, 3, 4, 6), X},
		{MakeGame(0, 1, 8, 4, 6, 7), O},
		{MakeGame(2, 0, 5, 1, 8), X},
		// diagonal wins
		{MakeGame(0, 1, 4, 1, 8), X},
		{MakeGame(0, 2, 1, 4, 8, 6), O},
	}

	for _, c := range cases {
		actual := c.input.Winner()
		Equal(t, actual, c.expected)
	}
}

func TestCanFindRandomMoves(t *testing.T) {
	cases := []struct {
		seed     int64
		input    Game
		expected int
	}{
		{1, MakeGame(0, 1, 2), 5},
		{1, MakeGame(), 5},
		{1, MakeGame(0, 1, 2, 3, 4, 5, 6, 7, 8), -1},
		{1, MakeGame(5), 6},
		{1, MakeGame(0, 1, 2, 3, 4, 5, 6, 7), 8},
	}

	for _, c := range cases {
		actual := c.input.GetRandomMove(c.seed)
		Equal(t, actual, c.expected)
	}
}

func TestCanFindBlockingMoves(t *testing.T) {
	cases := []struct {
		input    Game
		expected int
	}{
		// rows
		{MakeGame(0, 8, 2), 1},
		{MakeGame(3, 0, 4), 5},
		{MakeGame(7, 0, 8), 6},
		// columns
		{MakeGame(0, 1, 6), 3},
		{MakeGame(0, 1, 2, 4), 7},
		{MakeGame(2, 0, 8), 5},
		// diagonal
		{MakeGame(4, 1, 8), 0},
		{MakeGame(2, 5, 6), 4},
		// no block
		{MakeGame(), -1},
	}

	for _, c := range cases {
		actual := c.input.GetBlockingMove()
		Equal(t, actual, c.expected)
	}
}
