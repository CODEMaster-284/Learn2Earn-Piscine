package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	} else if n == 0 {
		z01.PrintRune('0')
		return
	} else {
		n = -n
	}
	printRecursive(n)
}

func printRecursive(n int) {
	if n <= -10 {
		printRecursive(n / 10)
	}

	digit := '0' - rune(n%10)
	z01.PrintRune(digit)
}
