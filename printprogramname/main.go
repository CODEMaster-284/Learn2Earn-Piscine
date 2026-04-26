package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[0]
	start := 0

	// Find last '/'
	for i := 0; i < len(arg); i++ {
		if arg[i] == '/' {
			start = i + 1
		}
	}

	// Print from last '/' to end
	for i := start; i < len(arg); i++ {
		z01.PrintRune(rune(arg[i]))
	}

	z01.PrintRune('\n')
}
