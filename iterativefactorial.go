package main

import "fmt"

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}

	result := 1

	for i := 1; i <= nb; i++ {
		result *= i

		if result < 0 {
			return 0
		}
	}

	return result
}

func main() {
	fmt.Print(IterativeFactorial(4))
}
