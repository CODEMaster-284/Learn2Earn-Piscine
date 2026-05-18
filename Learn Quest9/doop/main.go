package main

import "os"

const (
	Max int64 = 9223372036854775807
	Min int64 = -9223372036854775807 - 1
)

func printStr(s string) {
	os.Stdout.Write([]byte(s))
}

func printNbr(n int64) {
	if n == Min {
		printStr("-9223372036854775808")
		return
	}
	if n < 0 {
		printStr("-")
		n = -n
	}
	if n >= 10 {
		printNbr(n / 10)
	}
	os.Stdout.Write([]byte{byte(n%10 + '0')})
}

func atoi(s string) (int64, bool) {
	if len(s) == 0 {
		return 0, false
	}

	sign := int64(1)
	i := 0

	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}

	if i == len(s) {
		return 0, false
	}

	var n int64

	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}

		d := int64(s[i] - '0')

		if sign == 1 {
			if n > (Max-d)/10 {
				return 0, false
			}
		} else {
			if n > ((Max - d) / 10) {
				if !(n == Max/10 && d == 8) {
					return 0, false
				}
			}
		}

		n = n*10 + d
	}

	if sign == -1 {
		return -n, true
	}
	return n, true
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

	var r int64

	switch op {
	case "+":
		if (b > 0 && a > Max-b) || (b < 0 && a < Min-b) {
			return
		}
		r = a + b
	case "-":
		if (b < 0 && a > Max+b) || (b > 0 && a < Min+b) {
			return
		}
		r = a - b
	case "*":
		if a != 0 && b != 0 {
			if a == Min && b == -1 {
				return
			}
			if b == Min && a == -1 {
				return
			}
			r = a * b
			if r/b != a {
				return
			}
		} else {
			r = 0
		}
	case "/":
		if b == 0 {
			printStr("No division by 0\n")
			return
		}
		if a == Min && b == -1 {
			return
		}
		r = a / b
	case "%":
		if b == 0 {
			printStr("No modulo by 0\n")
			return
		}
		r = a % b
	default:
		return
	}

	printNbr(r)
	printStr("\n")
}
