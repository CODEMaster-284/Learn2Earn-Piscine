package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
		r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		z01.PrintRune('\n')
		return
	}

	// Join arguments with space
	str := ""
	for i := 0; i < len(args); i++ {
		str += args[i]
		if i != len(args)-1 {
			str += " "
		}
	}

	runes := []rune(str)

	// Step 1: collect vowels
	vowels := []rune{}
	for _, ch := range runes {
		if isVowel(ch) {
			vowels = append(vowels, ch)
		}
	}

	// Step 2: reverse vowels
	for i, j := 0, len(vowels)-1; i < j; i, j = i+1, j-1 {
		vowels[i], vowels[j] = vowels[j], vowels[i]
	}

	// Step 3: replace vowels in original string
	index := 0
	for i := 0; i < len(runes); i++ {
		if isVowel(runes[i]) {
			runes[i] = vowels[index]
			index++
		}
	}

	// Step 4: print result
	for _, ch := range runes {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
