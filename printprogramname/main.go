package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	start := 0

	for i, ch := range name {
		if ch == '/' {
			start = i + 1
		}
	}

	for _, ch := range name[start:] {
		z01.PrintRune(ch)
	}

	z01.PrintRune('\n')
}
