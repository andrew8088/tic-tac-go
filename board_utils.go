package main

import "fmt"

func FormatBoard(moves Moves) string {
	return fmt.Sprintf("%s", moves.ToBoard())
}

func FindWinner(moves Moves) string {
	return ""
}
