package main

import (
	"github.com/01-edu/z01"
)

func Rot14(s string) string {
	r := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			r += string((s[i]-'a'+14)%26 + 'a')
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			r += string((s[i]-'A'+14)%26 + 'A')
		} else {
			r += string(s[i])
		}
	}
	return r
}

func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
