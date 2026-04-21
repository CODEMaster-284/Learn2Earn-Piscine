package main

import "github.com/01-edu/z01"

func main() {
	// z01.PrintRune('H')
	// z01.PrintRune('e')
	// z01.PrintRune('l')
	// z01.PrintRune('l')
	// z01.PrintRune('0')

	str := "Hello"
	for _, char := range str {
		z01.PrintRune(char)
	}
}
