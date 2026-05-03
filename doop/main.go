package main

import (
	"fmt"
	"os"
)

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
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("No division by 0")
			return
		}
		fmt.Println(a / b)
	case "%":
		if b == 0 {
			fmt.Println("No modulo by 0")
			return
		}
		fmt.Println(a % b)
	default:
		return
	}
}
