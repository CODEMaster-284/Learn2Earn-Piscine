package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]

	// If no arguments → read from stdin
	if len(args) == 0 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	hasError := false

	for _, fileName := range args {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			hasError = true
			continue
		}

		io.Copy(os.Stdout, file)
		file.Close()
	}

	if hasError {
		os.Exit(1)
	}
}
