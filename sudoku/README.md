# Sudoku

This folder contains a Sudoku solver written in Go. The program expects a
9-row puzzle from the command line, validates the starting grid, solves it by
backtracking, and prints the finished board only when there is exactly one
solution.

## Files

- `main.go`
- `input.txt`

## Overall flow

The solver follows this pipeline:

1. Read the 9 puzzle rows from command-line arguments
2. Build the internal board representation
3. Validate the starting numbers as they are loaded
4. Solve the puzzle recursively
5. Count how many solutions exist
6. Print the solved board only if there is exactly one solution

If anything goes wrong, it prints:

`Error`

## `main.go` breakdown

### `main()`

This function controls the entire program.

Step by step:

1. `len(os.Args) != 10` checks that the program received exactly:
   - 1 program name
   - 9 Sudoku rows
2. If the argument count is wrong, print `Error`
3. Create `board := [9][9]int{}`
4. Call `fillBoard(os.Args[1:], &board)`
5. If the starting puzzle is invalid, print `Error`
6. Create:
   - `solutions := 0`
   - `answer := [9][9]int{}`
7. Call `solve(&board, &answer, &solutions)`
8. If the puzzle has zero or multiple solutions, print `Error`
9. Otherwise print the solved board with `printBoard(answer)`

So `main` is mostly orchestration: validate, solve, print.

### `fillBoard(args []string, board *[9][9]int) bool`

This function converts the 9 input strings into the board array.

Rules used while loading:

- each row must contain exactly 9 characters
- `.` means an empty cell, stored as `0`
- `'1'` to `'9'` are converted into integers
- every given number is checked immediately with `canPlace`

Why that validation matters:

- it prevents illegal starting boards
- it catches duplicate numbers in a row, column, or 3x3 box before solving even
  starts

If any row is malformed or any number breaks Sudoku rules, the function returns
`false`.

### `solve(board, answer, solutions)`

This is the recursive backtracking solver.

Its logic is:

1. If more than one solution has already been found, stop exploring
2. Find the next empty cell with `findEmpty`
3. If there is no empty cell:
   - a complete solution has been found
   - increase the solution count
   - copy the solved board into `answer`
   - return
4. Otherwise, try numbers `1` through `9`
5. For each number:
   - check if it is valid with `canPlace`
   - place it temporarily
   - recurse
   - undo the placement after returning

That "place, recurse, undo" pattern is the key backtracking idea.

### `findEmpty(board)`

This helper scans the board from top-left to bottom-right and returns:

- the row index
- the column index
- `true` if an empty cell was found

If no empty cell exists, it returns `false`, which means the puzzle is fully
filled.

### `canPlace(board, row, col, num)`

This is the rule checker for Sudoku placements.

It verifies that `num` does not already appear in:

- the same row
- the same column
- the same 3x3 subgrid

The 3x3 subgrid is located with:

- `startRow := row / 3 * 3`
- `startCol := col / 3 * 3`

That calculation snaps the current position back to the top-left corner of its
box.

If the number is safe in all three checks, the function returns `true`.

### `printBoard(board [9][9]int)`

This prints the solved grid.

Formatting details:

- numbers in the same row are separated by spaces
- each row is printed on its own line

So the final output is easy to read and matches the expected plain-text board
layout.

## `input.txt`

This file is not puzzle data by itself. It stores an example command showing
how to run the solver:

```text
go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
```

That example demonstrates:

- the required 9 command-line row inputs
- the use of `.` for blanks
- a quick way to inspect the output in the terminal

## Key algorithm idea

The heart of the folder is recursive backtracking:

1. choose an empty cell
2. try a candidate number
3. continue solving from that new state
4. back up if the choice leads to a dead end

This is a standard and effective way to solve Sudoku.

## Summary

This solver is built from small focused helpers:

- `fillBoard` loads and validates input
- `canPlace` enforces Sudoku rules
- `findEmpty` locates work to do
- `solve` performs the backtracking search
- `printBoard` formats the final answer

Together they form a complete command-line Sudoku solver.
