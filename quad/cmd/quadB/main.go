package main

import (
	"quad/internal/quadcmd"
	"quad/piscine"
)

func main() {
	x, y, ok := quadcmd.ParseSize()
	if !ok {
		return
	}

	piscine.QuadB(x, y)
}
