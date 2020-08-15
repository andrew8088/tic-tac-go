# tic-tac-go

A toy project to learn Go: super-simple tic-tac-toe game.

## Features

* accepts an unfinished game and returns the next move
* has different difficulty levels
* prints the game board

## Domain

* `Game [9]int` - defaults to all `-1`. Represents the current state of the
game by tracking all moves that have been played. Turns are played in order, so
all even indicies represent moves by `X` and all odd indicies represent moves
by `O`. The int represents the cell they made their move in:

```
0 | 1 | 2
3 | 4 | 5
6 | 7 | 8
```

* `Board [9] int` - Represents the board, where the items are the index of the
turns that have been played (terrible definition, to be improved).

* `Line struct{ indicies [3]int, moves [3]int }` - represents an individual
line in a game: could be a row, column, or diagonal.

## TODO

- [x] print boards of varying layouts
- [x] check games for a winner
- [x] play a random move
- [x] identify "win in one" layouts
- [x] play a strategic move
- [ ] deploy to vercel
- [ ] build a basic UI
