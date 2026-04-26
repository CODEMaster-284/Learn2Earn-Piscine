package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	baseLen := len(base)

	// Handle negative
	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}

	// Recursive conversion
	if nbr >= baseLen {
		PrintNbrBase(nbr/baseLen, base)
	}

	z01.PrintRune(rune(base[nbr%baseLen]))
}

// 🔍 Base validation function
func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	for i := 0; i < len(base); i++ {
		// Check + or -
		if base[i] == '+' || base[i] == '-' {
			return false
		}

		// Check duplicates
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}

	return true
}
