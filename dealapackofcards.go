package main

import (
	"fmt"
)

func DealAPackOfCards(deck []int) {
	p1 := deck[0:3]
	p2 := deck[3:6]
	p3 := deck[6:9]
	p4 := deck[9:12]

	fmt.Printf("Player 1: %d, %d, %d\n", p1[0], p1[1], p1[2])
	fmt.Printf("Player 2: %d, %d, %d\n", p2[0], p2[1], p2[2])
	fmt.Printf("Player 3: %d, %d, %d\n", p3[0], p3[1], p3[2])
	fmt.Printf("Player 4: %d, %d, %d\n", p4[0], p4[1], p4[2])
}

func main() {
	deck := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	DealAPackOfCards(deck)
}
