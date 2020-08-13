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
		{MakeGame(), Board{}},
		{MakeGame(0), Board{"x"}},
		{MakeGame(0, 7), Board{"x", "", "", "", "", "", "", "o", ""}},
		{MakeGame(0, 7, 1, 2), Board{"x", "x", "o", "", "", "", "", "o", ""}},
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
