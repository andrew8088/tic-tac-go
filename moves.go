package main

type Moves [9]int

func (m Moves) ToBoard() Board {
	board := Board{}

	for i, move := range m {
		// zero-value, no move
		if move == 0 {
			break
		}
		if i%2 == 0 {
			board[move-1] = "x"
		} else {
			board[move-1] = "o"
		}
	}
	return board
}
