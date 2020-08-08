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
		{MakeMoves(0, 3, 1, 4, 2), "x"},
		{MakeMoves(0, 3, 1, 4, 6, 5), "o"},
	}

	for _, c := range cases {
		actual := FindWinner(c.input)
		Equal(t, actual, c.expected)
	}
}
