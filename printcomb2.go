package main

import "github.com/01-edu/z01"

func main() {
	for first := 10; first <= 98; first++ {
		for second := first + 1; second <= 99; second++ {
			z01.PrintRune(rune('0' + first/10))
			z01.PrintRune(rune('0' + first%10))
			z01.PrintRune(' ')
			z01.PrintRune(rune('0' + second/10))
			z01.PrintRune(rune('0' + second%10))

			if first != 98 || second != 99 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
