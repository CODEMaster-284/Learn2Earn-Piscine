package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 || hasHelp(args) {
		printHelp()
		return
	}

	insert := ""
	order := false
	mainStr := ""

	for _, arg := range args {
		if startsWith(arg, "--insert=") {
			insert = arg[9:]
		} else if startsWith(arg, "-i=") {
			insert = arg[3:]
		} else if arg == "--order" || arg == "-o" {
			order = true
		} else {
			mainStr = arg
		}
	}

	result := mainStr + insert

	if order {
		result = sortString(result)
	}

	printStr(result)
	z01.PrintRune('\n')
}

func hasHelp(args []string) bool {
	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			return true
		}
	}
	return false
}

func printHelp() {
	printStr("--insert\n")
	printStr("  -i\n")
	printStr("\t This flag inserts the string into the string passed as argument.\n")
	printStr("--order\n")
	printStr("  -o\n")
	printStr("\t This flag will behave like a boolean, if it is called it will order the argument.\n")
}

func printStr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
}

func startsWith(s, prefix string) bool {
	if len(s) < len(prefix) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if s[i] != prefix[i] {
			return false
		}
	}
	return true
}

func sortString(s string) string {
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		for j := 0; j < len(runes)-1-i; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}

	return string(runes)
}
