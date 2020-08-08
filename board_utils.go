package main

func FormatBoard(game [9]int) string {

	if game[0] == 1 {
		return " x |   |   \n   |   |   \n   |   |   "
	}
	return "   |   |   \n   |   |   \n   |   |   "
}
