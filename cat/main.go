package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	for _, name := range args {
		file, err := os.Open(name)
		if err != nil {
			printStr("ERROR: " + err.Error() + "\n")
			os.Exit(1)
		}

		io.Copy(os.Stdout, file)
		file.Close()
	}
}
