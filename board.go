package main

import "fmt"

type Board [9]int

func (b Board) ToSymbols() [9]string {
	symbols := [9]string{}

	for i, move := range b {
		if move == -1 { // not a move
			continue
		}
		if move%2 == 0 {
			symbols[i] = "x"
		} else {
			symbols[i] = "o"
		}
	}

	return symbols
}

func (b Board) String() string {
	symbols := b.ToSymbols()
	return fmt.Sprintf(" %1s | %1s | %1s \n %1s | %1s | %1s \n %1s | %1s | %1s ", symbols[0], symbols[1], symbols[2], symbols[3], symbols[4], symbols[5], symbols[6], symbols[7], symbols[8])
}
