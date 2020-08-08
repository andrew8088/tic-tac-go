package main

import "testing"

func TestCanMakeMoves(t *testing.T) {
	cases := []struct {
		actual   Moves
		expected Moves
	}{
		{MakeMoves(), Moves{-1, -1, -1, -1, -1, -1, -1, -1, -1}},
		{MakeMoves(0), Moves{0, -1, -1, -1, -1, -1, -1, -1, -1}},
		{MakeMoves(0, 8, 1, 7), Moves{0, 8, 1, 7, -1, -1, -1, -1, -1}},
	}

	for _, c := range cases {
		Equal(t, c.actual, c.expected)
	}
}

func TestCanConvertMovesToBoard(t *testing.T) {
	cases := []struct {
		input    Moves
		expected Board
	}{
		{MakeMoves(), Board{}},
		{MakeMoves(0), Board{"x"}},
		{MakeMoves(0, 7), Board{"x", "", "", "", "", "", "", "o", ""}},
		{MakeMoves(0, 7, 1, 2), Board{"x", "x", "o", "", "", "", "", "o", ""}},
	}

	for _, c := range cases {
		actual := c.input.ToBoard()
		Equal(t, actual, c.expected)
	}
}

func TestCanConvertMovesToString(t *testing.T) {
	cases := []struct {
		input    Moves
		expected string
	}{
		{MakeMoves(), "[]"},
		{MakeMoves(1), "[1]"},
		{MakeMoves(1, 8), "[1, 8]"},
		{MakeMoves(1, 8, 2, 3), "[1, 8, 2, 3]"},
	}

	for _, c := range cases {
		actual := c.input.String()
		Equal(t, actual, c.expected)
	}
}
