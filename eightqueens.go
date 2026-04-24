package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int

	var solve func(col int)
	solve = func(col int) {
		if col == 8 {
			printBoard(board)
			return
		}

		for row := 1; row <= 8; row++ {
			if isSafe(board, col, row) {
				board[col] = row
				solve(col + 1)
			}
		}
	}

	solve(0)
}

func isSafe(board [8]int, col int, row int) bool {
	for prevCol := 0; prevCol < col; prevCol++ {
		prevRow := board[prevCol]

		if prevRow == row {
			return false
		}

		if abs(prevRow-row) == abs(prevCol-col) {
			return false
		}
	}

	return true
}

func printBoard(board [8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(board[i] + '0'))
	}
	z01.PrintRune('\n')
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
