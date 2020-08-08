package main

import "fmt"

type Board [9]string

func (b Board) String() string {
	return fmt.Sprintf("%3s|%3s|%3s\n%3s|%3s|%3s\n%3s|%3s|%3s", b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8])
}
