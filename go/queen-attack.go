package gocopilot

import (
	"fmt"
	"math"
)

/*
Convert this code to go
module QueenAttack

let create (position: int * int) =
    let x, y = position
    x > 0 && y > 0 && x < 8 && y < 8

let canAttack (queen1: int * int) (queen2: int * int) =
    let col1, row1 = queen1
    let col2, row2 = queen2
    col1 = col2 || row1 = row2 || abs (col1 - col2) = abs (row1 - row2)
*/

func queenAttack() {
	queen1 := []int{3, 4}
	queen2 := []int{0, 0}
	fmt.Println(queen1, queen2)
	fmt.Println(canAttack(queen1, queen2))
}

func canAttack(queen1, queen2 []int) bool {
	if queen1[0] == queen2[0] || queen1[1] == queen2[1] ||
		math.Abs(queen1[0] - queen2[0]) == math.Abs(queen1[1] - queen2[1]) {
		return true
	}
	return false
}
