# Quad

This folder contains a small Go module for drawing the five classic `quad`
patterns. The code is split into three parts so the drawing logic stays
separate from command-line parsing.

## Structure

- `main.go`
- `internal/quadcmd/args.go`
- `piscine/quadA.go`
- `piscine/quadB.go`
- `piscine/quadC.go`
- `piscine/quadD.go`
- `piscine/quadE.go`
- `cmd/quadA/main.go`
- `cmd/quadB/main.go`
- `cmd/quadC/main.go`
- `cmd/quadD/main.go`
- `cmd/quadE/main.go`

## How the project is organized

There are two entry styles in this folder:

- `main.go` is a simple demo runner. It calls all five quad functions with a
  fixed size of `5 x 3`.
- The files inside `cmd/quadA` to `cmd/quadE` are proper command entrypoints.
  Each one reads two command-line arguments, then runs exactly one quad
  function.

The real drawing logic lives in the `piscine` package.

## File breakdown

### `main.go`

This file imports `quad/piscine` and calls:

- `piscine.QuadA(5, 3)`
- `piscine.QuadB(5, 3)`
- `piscine.QuadC(5, 3)`
- `piscine.QuadD(5, 3)`
- `piscine.QuadE(5, 3)`

Its job is just to show sample output. It does not parse user input.

### `internal/quadcmd/args.go`

This file contains `ParseSize() (int, int, bool)`.

What it does:

1. Checks that exactly two arguments were passed after the program name.
2. Converts both arguments from strings to integers using `strconv.Atoi`.
3. Returns `(x, y, true)` when parsing works.
4. Returns `(0, 0, false)` when the input is missing or invalid.

This keeps argument handling out of the drawing files.

### `cmd/quadA/main.go` to `cmd/quadE/main.go`

Each of these files follows the same pattern:

1. Call `quadcmd.ParseSize()`.
2. Stop immediately if parsing failed.
3. Call the matching drawing function from the `piscine` package.

Example:

- `cmd/quadA/main.go` runs `piscine.QuadA(x, y)`
- `cmd/quadE/main.go` runs `piscine.QuadE(x, y)`

These files are intentionally thin wrappers.

## Drawing logic

All five `Quad*` functions work the same way at a high level:

1. Reject invalid sizes with:
   `if x <= 0 || y <= 0 { return }`
2. Loop through rows from `1` to `y`
3. Loop through columns from `1` to `x`
4. Decide what character belongs in the current position
5. Print the character with `z01.PrintRune`
6. Print a newline after each row

The most important idea is that each cell is classified as one of:

- a corner
- a horizontal border
- a vertical border
- an inside space

## Pattern-by-pattern explanation

### `piscine/quadA.go`

`QuadA` draws with:

- corners: `o`
- top and bottom edges: `-`
- left and right edges: `|`
- inside: space

The condition:

`(row == 1 || row == y) && (col == 1 || col == x)`

means "this cell is both on a top/bottom row and on a left/right column", so
it must be a corner.

### `piscine/quadB.go`

`QuadB` changes the border style:

- top-left and bottom-right corners: `/`
- top-right and bottom-left corners: `\`
- all non-corner border cells: `*`
- inside: space

This version uses more specific corner checks before the border checks.

### `piscine/quadC.go`

`QuadC` draws with letters:

- top corners: `A`
- bottom corners: `C`
- all other border cells: `B`
- inside: space

This means the shape visually distinguishes the top from the bottom.

### `piscine/quadD.go`

`QuadD` uses:

- top-left and bottom-left corners: `A`
- top-right and bottom-right corners: `C`
- all other border cells: `B`
- inside: space

Compared to `QuadC`, the left and right sides determine the corner letters.

### `piscine/quadE.go`

`QuadE` uses:

- top-left and bottom-right corners: `A`
- top-right and bottom-left corners: `C`
- all other border cells: `B`
- inside: space

This creates a diagonal symmetry in the corner markers.

## Core idea behind every quad

The program does not store the whole rectangle in memory first. Instead, it
decides the correct rune for each `(row, col)` position and prints it
immediately. That keeps the implementation simple and direct.

## Example mental model

For a size like `5 x 3`, the loops visit:

- row 1, columns 1 to 5
- row 2, columns 1 to 5
- row 3, columns 1 to 5

At each position the code asks:

- Is this a corner?
- Else, is it on the top or bottom border?
- Else, is it on the left or right border?
- Else, print a space.

That one decision tree is the heart of the whole folder.
