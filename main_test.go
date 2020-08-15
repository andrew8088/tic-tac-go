package main

import "testing"

func TestCanAddAMove(t *testing.T) {
	cases := []struct {
		seed     int64
		input    Game
		expected Game
	}{
		{1, MakeGame(0, 1, 2), MakeGame(0, 1, 2, 5)},
		{1, MakeGame(), MakeGame(5)},
		{1, MakeGame(0, 1, 2, 3, 4, 5, 6, 7, 8), MakeGame(0, 1, 2, 3, 4, 5, 6, 7, 8)},
		{1, MakeGame(5), MakeGame(5, 6)},
		{1, MakeGame(0, 1, 2, 3, 4, 5, 6, 7), MakeGame(0, 1, 2, 3, 4, 5, 6, 7, 8)},
	}

	for _, c := range cases {
		actual := PlayRandomMove(c.seed, c.input)
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
	}

	for _, c := range cases {
		actual := FindBlockingMove(c.input)
		Equal(t, actual, c.expected)
	}
}
