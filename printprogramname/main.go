package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Get only the program name without the path
	name := filepath.Base(os.Args[0])
	fmt.Println(name)
}
