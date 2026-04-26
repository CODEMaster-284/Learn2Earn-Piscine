package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Store digits count (0–9)
	var digits [10]int

	// Extract digits
	for n > 0 {
		d := n % 10
		digits[d]++
		n /= 10
	}

	// Print digits in ascending order
	for i := 0; i <= 9; i++ {
		for digits[i] > 0 {
			z01.PrintRune(rune(i + '0'))
			digits[i]--
		}
	}
}
