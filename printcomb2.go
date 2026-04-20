package main

import "github.com/01-edu/z01"

func main() {
	for a := '0'; a <= '9'; a++ {
		for b := '0'; b <= '9'; b++ {
			for c := a; c <= '9'; c++ {
				start := '0'
				if c == a {
					start = b + 1
				}
				for d := start; d <= '9'; d++ {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(' ')
					z01.PrintRune(c)
					z01.PrintRune(d)

					if !(a == '9' && b == '8' && c == '9' && d == '9') {
						z01.PrintRune(',')
						z01.PrintRune(' ')

					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
