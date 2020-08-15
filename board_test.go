package main

import "testing"

func TestCanFormatBoards(t *testing.T) {
	cases := []struct {
		input    Game
		expected string
	}{
		{MakeGame(), "   |   |   \n   |   |   \n   |   |   "},
		{MakeGame(0), " x |   |   \n   |   |   \n   |   |   "},
		{MakeGame(0, 7), " x |   |   \n   |   |   \n   | o |   "},
		{MakeGame(0, 7, 1, 2), " x | x | o \n   |   |   \n   | o |   "},
	}

	for _, c := range cases {
		actual := c.input.ToBoard().String()
		Equal(t, actual, c.expected)
	}
}
