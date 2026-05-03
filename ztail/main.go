package main

import (
	"fmt"
	"os"
)

func atoi(s string) int {
	n := 0
	for _, r := range s {
		if r >= '0' && r <= '9' {
			n = n*10 + int(r-'0')
		}
	}
	return n
}

func main() {
	n := atoi(os.Args[2])
	files := os.Args[3:]
	hasError := false

	for i, name := range files {
		data, err := os.ReadFile(name)
		if err != nil {
			fmt.Printf("%v\n", err)
			hasError = true
			continue
		}

		if len(files) > 1 {
			if i > 0 {
				fmt.Printf("\n")
			}
			fmt.Printf("==> %s <==\n", name)
		}

		start := len(data) - n
		if start < 0 {
			start = 0
		}

		fmt.Printf("%s", data[start:])
	}

	if hasError {
		os.Exit(1)
	}
}
