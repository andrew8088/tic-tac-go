package main

import (
	"testing"
)

func TestCanFormatBoards(t *testing.T) {
	cases := []struct {
		input    [9]int
		expected string
	}{
		{[9]int{}, "   |   |   \n   |   |   \n   |   |   "},
		{[9]int{1}, " x |   |   \n   |   |   \n   |   |   "},
	}

	for _, c := range cases {
		actual := FormatBoard(c.input)

		if actual != c.expected {
			t.Errorf("actual:\n%s\nexpected:\n%s", actual, c.expected)
		}

	}

}
