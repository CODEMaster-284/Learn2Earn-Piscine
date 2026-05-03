package main

import "fmt"

// Define the struct
type point struct {
	x int
	y int
}

// Function to set values
func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
