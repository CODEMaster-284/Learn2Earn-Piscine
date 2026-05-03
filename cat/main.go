package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printError(err error) {
	msg := "ERROR: " + err.Error() + "\n"
	for _, r := range msg {
		z01.PrintRune(r)
	}
}

func main() {
	args := os.Args[1:]

	// No arguments → stdin
	if len(args) == 0 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	hasError := false

	for _, name := range args {
		file, err := os.Open(name)
		if err != nil {
			printError(err)
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
