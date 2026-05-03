package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args

	// No argument
	if len(args) == 1 {
		fmt.Println("File name missing")
		return
	}

	// More than one argument
	if len(args) > 2 {
		fmt.Println("Too many arguments")
		return
	}

	// Exactly one argument
	file, err := os.Open(args[1])
	if err != nil {
		return
	}
	defer file.Close()

	// Read and display file content
	io.Copy(os.Stdout, file)
}
