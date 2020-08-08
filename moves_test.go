package main

import "testing"

func TestCanConvertMovesToBoard(t *testing.T) {
	cases := []struct {
		input    Moves
		expected Board
	}{
		{Moves{}, Board{}},
		{Moves{1}, Board{"x"}},
		{Moves{1, 8}, Board{"x", "", "", "", "", "", "", "o", ""}},
		{Moves{1, 8, 2, 3}, Board{"x", "x", "o", "", "", "", "", "o", ""}},
	}

	for _, c := range cases {
		actual := c.input.ToBoard()

		if actual != c.expected {
			t.Errorf("\nactual:\n%s\nexpected:\n%s", actual, c.expected)
		}
	}
}

func TestCanConvertMovesToString(t *testing.T) {
	cases := []struct {
		input    Moves
		expected string
	}{
		{Moves{}, "[]"},
		{Moves{1}, "[1]"},
		{Moves{1, 8}, "[1, 8]"},
		{Moves{1, 8, 2, 3}, "[1, 8, 2, 3]"},
	}

	for _, c := range cases {
		actual := c.input.String()

		if actual != c.expected {
			t.Errorf("\nactual:\n%s\nexpected:\n%s", actual, c.expected)
		}
	}
}
