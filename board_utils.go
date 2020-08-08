package main

import "fmt"

func FormatBoard(moves Moves) string {
	return fmt.Sprintf("%s", moves.ToBoard())
}
