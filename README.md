# Piscine Go Repo Guide

This repository is organized to help you move through the Learn2Earn / Piscine
Go work in a cleaner way. Instead of keeping everything flat in one directory,
the exercises are grouped into quest folders and a few standalone mini-project
folders.

This README is meant to help you understand:

- how the root folder is arranged
- where each type of exercise lives
- how to navigate folders with spaces in their names
- which folders are standalone Go projects

## Root structure

These are the main folders you should care about from the root:

| Folder | Purpose |
| --- | --- |
| `Learn Quest1` | Shell tasks, basic command-line exercises, file/permission tasks |
| `Learn Quest2` | Early Go printing exercises like alphabet, digits, combinations |
| `Learn Quest3` | Small utility functions like `atoi`, `strlen`, `swap`, `divmod` |
| `Learn Quest4` | Recursion, factorials, powers, Fibonacci, primes, square root |
| `Learn Quest5` | String and rune work, casing, joins, bases, text parsing |
| `Learn Quest6` | Command-line Go programs like `printparams`, `sortparams`, `flags` |
| `Learn Quest7` | Slices, splitting, concatenation, base conversion, helper logic |
| `Learn Quest8` | More complete Go programs like `boolean`, `cat`, `displayfile`, `ztail` |
| `Learn Quest9` | Higher-order functions and sorting helpers like `Map`, `Any`, `Doop` |
| `Hackathon` | Miscellaneous challenge-style Go exercises |
| `Learn Quest11` | Linked list exercises |
| `Quest 12` | Binary tree exercises |
| `quad` | Standalone quad drawing project with multiple commands |
| `quadchecker` | Standalone quad recognition/checker project |
| `sudoku` | Standalone Sudoku solver project |
| `the-final-cl-test` | Preserved extra/testing-related folder from the original repo |

You will also see some repo-management files and folders:

| Item | Meaning |
| --- | --- |
| `.git` | Git metadata for the main repository |
| `.qodo` | Tooling/workflow metadata |
| `go.mod` | Root Go module definition |
| `go.sum` | Root Go dependency checksums |

If you notice a stray single file at the root that does not match the structure
above, treat it as legacy or leftover material, not part of the main guided
navigation.

## How to navigate

Some folder names contain spaces, so you should quote them when changing
directories.

### In Git Bash

```bash
cd "Learn Quest1"
cd "Learn Quest6"
cd "Quest 12"
cd "Hackathon"
cd "quad"
cd "quadchecker"
cd "sudoku"
```

To go back up one level:

```bash
cd ..
```

To see what is inside your current folder:

```bash
ls
```

### In PowerShell

```powershell
cd "Learn Quest1"
cd "Learn Quest6"
cd "Quest 12"
Get-ChildItem
```

## Suggested learning path

If someone is new to the repo, this is the easiest reading order:

1. Start in `Learn Quest1` for shell basics.
2. Move to `Learn Quest2` and `Learn Quest3` for simple Go fundamentals.
3. Continue to `Learn Quest4` and `Learn Quest5` for recursion, math, strings, and runes.
4. Use `Learn Quest6` to practice command-line programs.
5. Visit `Learn Quest7` to get comfortable with slices and string manipulation.
6. Go to `Learn Quest8` and `Learn Quest9` for slightly larger programs and function-based patterns.
7. Move into `Learn Quest11` and `Quest 12` for linked lists and binary trees.
8. Explore `quad`, `quadchecker`, and `sudoku` as separate small projects.

## What each major area contains

### `Learn Quest1`

This folder mostly contains shell tasks and filesystem-related exercises.
Examples include:

- permission tasks
- small `.sh` scripts
- basic command-line logic

### `Learn Quest2` to `Learn Quest5`

These folders are the core Go practice path.

- `Learn Quest2` focuses on output and simple loops
- `Learn Quest3` focuses on tiny utility functions
- `Learn Quest4` focuses on recursion and number problems
- `Learn Quest5` focuses on strings, runes, text formatting, and parsing

### `Learn Quest6` to `Learn Quest9`

These folders move from single functions into more program-style exercises.

- `Learn Quest6` contains command-based programs in their own subfolders
- `Learn Quest7` focuses on slices, splits, joins, and conversion helpers
- `Learn Quest8` contains utility programs like `cat` and `displayfile`
- `Learn Quest9` contains function-driven patterns like `Map`, `Any`, and sorting

### `Hackathon`

This folder groups mixed challenge tasks that do not fit as neatly into the
main quest progression. It is good for extra practice and revision.

### `Learn Quest11`

This is the linked list section. If you are looking for:

- `ListPushFront`
- `ListPushBack`
- `ListSize`
- `ListRemoveIf`

this is the right folder.

### `Quest 12`

This is the binary tree section. If you are looking for:

- `BTreeInsertData`
- `BTreeMax`
- `BTreeMin`
- traversal helpers
- delete/transplant helpers

this is the right folder.

## Standalone project folders

These three are separate mini-projects, not just single quest exercise dumps.

### `quad`

This is a full quad drawing module.

Important parts:

- `main.go` shows sample output
- `cmd/quadA` to `cmd/quadE` are command entrypoints
- `internal/quadcmd` handles size parsing
- `piscine` contains the real drawing logic

### `quadchecker`

This project reads quad text from standard input and tries to identify which
quad pattern produced it. It is useful as a checker/recognizer companion to the
`quad` project.

### `sudoku`

This is a Sudoku solver project. It reads a 9-line puzzle from command-line
arguments, validates it, solves it, and prints the solved board if the solution
is unique.

## Folder naming note

You may notice that most quest folders use names like:

- `Learn Quest1`
- `Learn Quest6`
- `Quest 12`

Those names are intentional and match the organized structure used for this
repo. Because some names contain spaces, always remember to use quotes in shell
commands.

## Quick examples

### Open the shell tasks

```bash
cd "Learn Quest1"
ls
```

### Open linked list exercises

```bash
cd "Learn Quest11"
ls
```

### Open the Sudoku project

```bash
cd sudoku
ls
```

### Return to the root from inside a project

```bash
cd ..
```

If you are deeper inside a nested folder, keep going back up:

```bash
cd ..
cd ..
```

## Repo usage advice

- Use the root folder when you want the full organized overview.
- Enter a quest folder when you want to focus on one learning stage.
- Enter `quad`, `quadchecker`, or `sudoku` when you want to work on those as
  standalone projects.
- Be careful with `git add .` when you are testing temporary files inside a
  project. It is usually safer to add only the files you actually changed.

## Final orientation

Think of the repo like this:

- the `Learn Quest*` folders are the main curriculum path
- `Hackathon`, `Learn Quest11`, and `Quest 12` are later or side sections
- `quad`, `quadchecker`, and `sudoku` are self-contained project folders
- the root folder is your map, not the place where you should keep hunting for
  every exercise manually

If you follow the structure from this README, it becomes much easier to know
where you are and where to go next.
