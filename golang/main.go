// 8 queens problem generalized to N queens
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type board []int

var nQueens = flag.Int("n", 8, "Number of queens to place (implies size of board)")
var loopFlag = flag.Bool("loop", false, "loop from n=1 ...")

func main() {
	flag.Parse()
	if !*loopFlag {
		solveForN(*nQueens)
		return
	}
	// infinite loop
	for n := 1; ; n++ {
		solveForN(n)
	}
}

// solveForN: solve N queens problem
func solveForN(n int) {
	solutions := solve(n)
	fmt.Printf("Queens: %d Solutions: %d\n\n", n, len(solutions))
}

// solve 8 queens problem, generalized to N queens
func solve(nQueens int) []board {
	stop := startTimer(fmt.Sprintf("%d Queens", nQueens))
	defer stop()
	// place initial queen
	boards := make([]board, nQueens)
	for i := 0; i < nQueens; i++ {
		boards[i] = []int{i}
	}
	// place remaining nQueens-1 queens
	for i := 1; i < nQueens; i++ {
		var grow []board
		for _, b := range boards {
			expanded := expand(b, nQueens)
			grow = append(grow, expanded...)
		}
		boards = grow
	}
	return boards
}

// startTimer: start a timer and return a function to stop the timer
func startTimer(name string) func() {
	t := time.Now()
	log.Println(name, "started")
	return func() {
		d := time.Now().Sub(t)
		log.Println(name, "took", d)
	}
}

// place a queen in all valid positions of next column
func expand(b board, nQueens int) []board {
	var boards []board
	for i := 0; i < nQueens; i++ {
		if isValid(b, i) {
			boards = append(boards, newBoard(b, i))
		}
	}
	return boards
}

// isValid: check if adding x is valid position to add to b
func isValid(b board, x int) bool {
	return isValidHorizontal(b, x) && isValidDiagonal(b, x)
}

// newBoard: create new board by appending x
func newBoard(b []int, x int) []int {
	new := make([]int, len(b)+1)
	for i := 0; i < len(b); i++ {
		new[i] = b[i]
	}
	new[len(new)-1] = x
	return new
}

// check that last element is valid horizontally
func isValidHorizontal(b []int, x int) bool {
	for i := 0; i < len(b); i++ {
		if b[i] == x {
			return false
		}
	}
	return true
}

// check that last element is valid diagonally
func isValidDiagonal(b []int, x int) bool {
	for i := 0; i < len(b); i++ {
		delta := len(b) - i
		if b[i] == x-delta || b[i] == x+delta {
			return false
		}
	}
	return true
}

// printSolutions: print solutions to stdout
func printSolutions(s []board) {
	for _, b := range s {
		b.write(os.Stdout)
	}
	fmt.Printf("Total solutions: %d\n", len(s))
}

// write b to w
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
