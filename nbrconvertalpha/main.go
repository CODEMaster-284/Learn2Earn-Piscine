package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	upper := false
	start := 0

	if len(args) == 0 {
		return
	}

	if args[0] == "--upper" {
		upper = true
		start = 1
	}

	for i := start; i < len(args); i++ {
		n := atoi(args[i])

		if n >= 1 && n <= 26 {
			if upper {
				z01.PrintRune(rune('A' + n - 1))
			} else {
				z01.PrintRune(rune('a' + n - 1))
			}
		} else {
			z01.PrintRune(' ')
		}
	}

	z01.PrintRune('\n')
}

func atoi(s string) int {
	n := 0

	if s == "" {
		return 0
	}

	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return 0
		}
		n = n*10 + int(ch-'0')
	}

	return n
}
