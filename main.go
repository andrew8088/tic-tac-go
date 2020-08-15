package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("tttapi")
}

func PlayRandomMove(seed int64, game Game) Game {
	rand.Seed(seed)

	if !contains(game[:], -1) {
		return game // game is complete
	}

	move := rand.Intn(9)

	for contains(game[:], move) {
		move = rand.Intn(9)
	}

	for i, m := range game {
		if m == -1 {
			game[i] = move
			break
		}
	}

	return game
}

func FindBlockingMove(game Game) int {
	lines := game.Lines()

	for _, line := range lines {
		mod1 := line.indices[0] % 2
		mod2 := line.indices[1] % 2
		mod3 := line.indices[2] % 2

		if all(line.indices[:], getIsNum(-1)) {
			continue
		}

		if mod1 == mod2 && mod3 == -1 {
			return line.moves[2]
		}

		if mod1 == mod3 && mod2 == -1 {
			return line.moves[1]

		}

		if mod2 == mod3 && mod1 == -1 {
			return line.moves[0]
		}
	}

	return -1
}

func getIsNum(num1 int) func(int) bool {
	return func(num2 int) bool {
		return num1 == num2
	}
}

type predicate func(int) bool

func any(arr []int, predicate predicate) bool {
	for _, m := range arr {
		if predicate(m) {
			return true
		}
	}
	return false
}

func all(arr []int, predicate predicate) bool {
	for _, m := range arr {
		if !predicate(m) {
			return false
		}
	}
	return true
}

func contains(arr []int, item int) bool {
	return any(arr, getIsNum(item))
}
