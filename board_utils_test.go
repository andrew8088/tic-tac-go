package main

import (
	"testing"
)

func TestCanFormatBoards(t *testing.T) {
	cases := []struct {
		input    Moves
		expected string
	}{
		{Moves{}, "   |   |   \n   |   |   \n   |   |   "},
		{Moves{1}, " x |   |   \n   |   |   \n   |   |   "},
		{Moves{1, 8}, " x |   |   \n   |   |   \n   | o |   "},
		{Moves{1, 8, 2, 3}, " x | x | o \n   |   |   \n   | o |   "},
	}

	for _, c := range cases {
		actual := FormatBoard(c.input)

		if actual != c.expected {
			t.Errorf("\nactual:\n%s\nexpected:\n%s", actual, c.expected)
		}

	}

}
