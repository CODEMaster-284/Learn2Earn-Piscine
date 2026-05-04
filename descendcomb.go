package main

import (
	"fmt"
)

func printDescendingComb() {
	first := true
	// Outer loop for the first two-digit number (99 to 01)
	for i := 99; i >= 1; i-- {
		// Inner loop for the second two-digit number (strictly smaller than i)
		for j := i - 1; j >= 0; j-- {
			// Print separator for all except the very first combination
			if !first {
				fmt.Print(", ")
			}
			// Print the numbers formatted as two digits
			fmt.Printf("%02d %02d", i, j)
			first = false
		}
	}
	fmt.Println()
}

func main() {
	printDescendingComb()
}
