package main

import "fmt"

func FormatBoard(game [9]int) string {
    board := [9]string{}

    for i, move := range game {
        // zero-value, no move
        if move == 0 {
            break
        }
        if i % 2 == 0 {
            board[move - 1] = "x"
        } else {
            board[move - 1] = "o"
        }
    }

    return fmt.Sprintf(" %1s | %1s | %1s \n %1s | %1s | %1s \n %1s | %1s | %1s ", board[0], board[1], board[2], board[3], board[4], board[5], board[6], board[7], board[8])
}
