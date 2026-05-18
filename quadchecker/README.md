# Quad Checker

> [!WARNING]
> This repository content is for learning and personal practice only.
> Do not use it during checkpoints, recoding sessions, exams, or any other
> graded or evaluated activity during the piscine.
> If you choose to misuse it, you alone are responsible for the consequences.

This folder contains a standalone quad recognizer. Instead of drawing one
pattern directly, it reads a shape from standard input, figures out its width
and height, regenerates all known quad patterns for that size, and reports
which one matches.

## Files

- `main.go`

## What the program does

The checker solves this problem:

1. Read an ASCII rectangle from `stdin`
2. Measure its dimensions
3. Generate what `quadA`, `quadB`, `quadC`, `quadD`, and `quadE` would look
   like for those dimensions
4. Compare the input against each generated result
5. Print the matching quad name and size

If nothing matches, it prints:

`Not a quad function`

If more than one quad matches, it prints all matches separated by ` || `.

## `main.go` breakdown

### `quadA`, `quadB`, `quadC`, `quadD`, `quadE`

These five functions mirror the drawing rules from the quad exercise, but with
one important difference:

- they build and return a `string`
- they do not print directly with `z01.PrintRune`

That design makes sense for a checker, because strings are easy to compare with
the input.

Each function:

1. Returns `""` when `x` or `y` is invalid
2. Uses nested loops over rows and columns
3. Appends the correct character to `result`
4. Appends `\n` after each row
5. Returns the final string

So these functions are "generators" for expected output.

### `getDimensions(input string) (int, int)`

This helper measures the incoming shape.

How it works:

1. If the input is empty, it returns `0, 0`
2. It scans the string byte by byte
3. `currentWidth` counts characters until a newline
4. Every newline increases `height`
5. The first completed line becomes the stored `width`

The function assumes the input ends rows with newline characters, which matches
the output style of the quad generators.

### `main()`

This is the control center of the checker.

Step by step:

1. `io.ReadAll(os.Stdin)` reads the full incoming shape
2. The data is converted to a string
3. `getDimensions` computes `x` and `y`
4. An empty `matches` slice is created
5. The program compares the input to:
   - `quadA(x, y)`
   - `quadB(x, y)`
   - `quadC(x, y)`
   - `quadD(x, y)`
   - `quadE(x, y)`
6. Every successful match adds a formatted message like:
   `[quadA] [5] [3]`
7. If there are no matches, print `Not a quad function`
8. Otherwise print every match, joined by ` || `

## Why multiple matches are possible

For some very small sizes, different quad functions can produce the same final
output. For example, a `1 x 1` shape may collapse to a single corner character,
which can make patterns ambiguous.

That is why the checker stores matches in a slice instead of stopping at the
first one.

## Important design choice

This folder separates two ideas clearly:

- generator functions create the expected shapes
- `main` handles input, comparison, and reporting

That makes the checker easy to read:

- if you want pattern rules, read the `quad*` functions
- if you want matching logic, read `main`

## Summary

The code works like a small verifier:

- read
- measure
- regenerate
- compare
- report

That is the central flow of the whole folder.
