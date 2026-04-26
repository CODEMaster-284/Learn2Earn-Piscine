package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !validBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	if nbr < 0 {
		z01.PrintRune('-')
		printNegative(nbr, base, len(base))
	} else {
		printPositive(nbr, base, len(base))
	}
}

func validBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			return false
		}

		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}

	return true
}

func printPositive(nbr int, base string, baseLen int) {
	if nbr >= baseLen {
		printPositive(nbr/baseLen, base, baseLen)
	}

	z01.PrintRune(rune(base[nbr%baseLen]))
}

func printNegative(nbr int, base string, baseLen int) {
	if nbr <= -baseLen {
		printNegative(nbr/baseLen, base, baseLen)
	}

	rem := nbr % baseLen
	if rem < 0 {
		rem = -rem
	}

	z01.PrintRune(rune(base[rem]))
}
