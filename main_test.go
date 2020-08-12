package main

import "testing"

func TestCanFormatBoards(t *testing.T) {
	cases := []struct {
		input    Moves
		expected string
	}{
		{MakeMoves(), "   |   |   \n   |   |   \n   |   |   "},
		{MakeMoves(0), " x |   |   \n   |   |   \n   |   |   "},
		{MakeMoves(0, 7), " x |   |   \n   |   |   \n   | o |   "},
		{MakeMoves(0, 7, 1, 2), " x | x | o \n   |   |   \n   | o |   "},
	}

	for _, c := range cases {
		actual := FormatBoard(c.input)
		Equal(t, actual, c.expected)
	}
}

func TestCanDetermineTheWinner(t *testing.T) {
	cases := []struct {
		input    Moves
		expected string
	}{
		{MakeMoves(), ""},
		// row wins
		{MakeMoves(0, 3, 1, 4, 2), "x"},
		{MakeMoves(0, 3, 1, 4, 6, 5), "o"},
		{MakeMoves(8, 3, 7, 4, 6), "x"},
		// columns wins
		{MakeMoves(0, 1, 3, 4, 6), "x"},
		{MakeMoves(0, 1, 8, 4, 6, 7), "o"},
		{MakeMoves(2, 0, 5, 1, 8), "x"},
		// diagonal wins
		{MakeMoves(0, 1, 4, 1, 8), "x"},
		{MakeMoves(0, 2, 1, 4, 8, 6), "o"},
	}

	for _, c := range cases {
		actual := FindWinner(c.input)
		Equal(t, actual, c.expected)
	}
}

func TestCanAddAMove(t *testing.T) {
	cases := []struct {
		seed     int64
		input    Moves
		expected Moves
	}{
		{1, MakeMoves(0, 1, 2), MakeMoves(0, 1, 2, 5)},
		{1, MakeMoves(), MakeMoves(5)},
		{1, MakeMoves(0, 1, 2, 3, 4, 5, 6, 7, 8), MakeMoves(0, 1, 2, 3, 4, 5, 6, 7, 8)},
	}

	for _, c := range cases {
		actual := PlayRandomMove(c.seed, c.input)
		Equal(t, actual, c.expected)
	}
}
