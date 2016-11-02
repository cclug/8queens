// 8 queens problem
package main

import (
	"fmt"
	"io"
	"os"
)

type board []int

const size = 8

func main() {
	solutions := solve()
	for _, b := range solutions {
		b.write(os.Stdout)
	}
	fmt.Printf("\ntotal solutions: %d\n", len(solutions))
}

// return all solutions to 8 queens problem
func solve() []board {
	// place initial queen
	boards := make([]board, size)
	for i := 0; i < size; i++ {
		boards[i] = []int{i}
	}
	// place remaining size-1 queens
	for i := 1; i < size; i++ {
		var grow []board
		for _, b := range boards {
			expanded := expand(b)
			grow = append(grow, expanded...)
		}
		boards = grow
	}
	//	fmt.Printf("%v\n", boards)
	return boards
}

func (b board) write(w io.Writer) {
	var slice []int = b
	fmt.Fprintf(w, "(")
	for i, x := range slice {
		fmt.Fprintf(w, "%d", x)
		if i != len(slice)-1 {
			fmt.Fprintf(w, " ")
		}
	}
	fmt.Fprintf(w, ")\n")
}

// place a queen in all valid positions of next column
func expand(b board) []board {
	var valid []board
	for i := 0; i < size; i++ {
		try := append(b, i)
		if isValid(try) {
			valid = append(valid, copy(try))
		}
	}
	return valid
}

func isValid(b board) bool {
	return isValidHorizontal(b) && isValidDiagonal(b)
}

// check that last element is valid horizontally
func isValidHorizontal(b []int) bool {
	v := b[len(b)-1]
	for i := 0; i < len(b)-1; i++ {
		if b[i] == v {
			return false
		}
	}
	return true
}

// check that last element is valid diagonally
func isValidDiagonal(b []int) bool {
	v := b[len(b)-1]
	for i := 0; i < len(b)-1; i++ {
		delta := len(b) - 1 - i
		if b[i] == v-delta || b[i] == v+delta {
			return false
		}
	}
	return true
}

func copy(xs []int) []int {
	cp := make([]int, len(xs))
	for i := range xs {
		cp[i] = xs[i]
	}
	return cp
}
