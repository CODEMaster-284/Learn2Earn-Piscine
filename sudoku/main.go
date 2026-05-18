package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}
	board := [9][9]int{}


	if !fillBoard(os.Args[1:], &board) {
		fmt.Println("Error")
		return
	}

	solutions := 0
	answer := [9][9]int{}


	solve(&board, &answer, &solutions)

	if solutions != 1 {
		fmt.Println("Error")
		return
	}

	printBoard(answer)
}


func fillBoard(args []string, board *[9][9]int) bool {
	for row := 0; row < 9; row++ {
		if len(args[row]) != 9 {
			return false
		}
		for col := 0; col < 9; col++ {
			ch := args[row][col]

			if ch == '.' {
				board[row][col] = 0
			} else if ch >= '1' && ch <= '9' {
		
				num := int(ch - '0')
				if !canPlace(board, row, col, num) {
					return false
				}
				// if the number is allowed place it on the board
				board[row][col] = num
			} else {
				return false
			}
		}
	}

	return true
}


func solve(board *[9][9]int, answer *[9][9]int, solutions *int) {
	// if more than one solution is already found stop
	if *solutions > 1 {
		return
	}

	// find the empty cell
	row, col, found := findEmpty(board)

	if !found {
		*solutions++

		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				answer[i][j] = board[i][j]
			}
		}

		// stop the solving path
		return
	}

	for num := 1; num <= 9; num++ {
		// check if the number is in that valid position
		if canPlace(board, row, col, num) {
			// place the number temporarily
			board[row][col] = num

		
			solve(board, answer, solutions)

		
			board[row][col] = 0
		}
	}
}

// searches for an empty cell
func findEmpty(board *[9][9]int) (int, int, bool) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == 0 {
				return row, col, true
			}
		}
	}

	// if no empty cell is found the board is complete
	return 0, 0, false
}


func canPlace(board *[9][9]int, row int, col int, num int) bool {
	for i := 0; i < 9; i++ {
		// board[row][i] same row different columns
		if board[row][i] == num {
			return false
		}
		// board[i][col] came column different row
		if board[i][col] == num {
			return false
		}
		// if the number exists already it would return false
	}

	// this just find the position of which 3 by 3 group that you are
	startRow := row / 3 * 3
	startCol := col / 3 * 3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return false
			}
		}
	}
	// if the number is not in the row,column,3 by 3 box it should return true
	return true
}

func printBoard(board [9][9]int) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			// print space before every number
			if col > 0 {
				fmt.Print(" ")
			}
			// print the numbers
			fmt.Print(board[row][col])
		}
		// move to the next line
		fmt.Println()
	}
}
