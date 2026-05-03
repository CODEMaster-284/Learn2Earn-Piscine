package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printNbr(n int) {
	if n == -2147483648 {
		printStr("-2147483648")
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n >= 10 {
		printNbr(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}

func atoi(s string) (int, bool) {
	n := 0
	sign := 1

	if len(s) == 0 {
		return 0, false
	}

	i := 0
	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}

	if i >= len(s) {
		return 0, false
	}

	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		n = n*10 + int(s[i]-'0')
	}

	return n * sign, true
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, ok1 := atoi(os.Args[1])
	op := os.Args[2]
	b, ok2 := atoi(os.Args[3])

	if !ok1 || !ok2 {
		return
	}

	switch op {
	case "+":
		printNbr(a + b)
	case "-":
		printNbr(a - b)
	case "*":
		printNbr(a * b)
	case "/":
		if b == 0 {
			printStr("No division by 0")
			z01.PrintRune('\n')
			return
		}
		printNbr(a / b)
	case "%":
		if b == 0 {
			printStr("No modulo by 0")
			z01.PrintRune('\n')
			return
		}
		printNbr(a % b)
	default:
		return
	}

	z01.PrintRune('\n')
}
