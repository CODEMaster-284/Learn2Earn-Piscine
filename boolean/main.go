package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

// Function to check if number is even
func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	}
	return false
}

func main() {
	argCount := len(os.Args) - 1

	if isEven(argCount) {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}