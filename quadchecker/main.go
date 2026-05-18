package main

import (
	"fmt"
	"io"
	"os"
)

func quadA(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}

	result := ""

	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if (row == 1 || row == y) && (col == 1 || col == x) {
				result += "o"
			} else if row == 1 || row == y {
				result += "-"
			} else if col == 1 || col == x {
				result += "|"
			} else {
				result += " "
			}
		}
		result += "\n"
	}

	return result
}

func quadB(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}

	result := ""

	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && col == 1 {
				result += "/"
			} else if row == 1 && col == x {
				result += "\\"
			} else if row == y && col == 1 {
				result += "\\"
			} else if row == y && col == x {
				result += "/"
			} else if row == 1 || row == y || col == 1 || col == x {
				result += "*"
			} else {
				result += " "
			}
		}
		result += "\n"
	}

	return result
}

func quadC(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}

	result := ""

	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && (col == 1 || col == x) {
				result += "A"
			} else if row == y && (col == 1 || col == x) {
				result += "C"
			} else if row == 1 || row == y || col == 1 || col == x {
				result += "B"
			} else {
				result += " "
			}
		}
		result += "\n"
	}

	return result
}

func quadD(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}

	result := ""

	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && col == 1 {
				result += "A"
			} else if row == 1 && col == x {
				result += "C"
			} else if row == y && col == 1 {
				result += "A"
			} else if row == y && col == x {
				result += "C"
			} else if row == 1 || row == y || col == 1 || col == x {
				result += "B"
			} else {
				result += " "
			}
		}
		result += "\n"
	}

	return result
}

func quadE(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}

	result := ""

	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && col == 1 {
				result += "A"
			} else if row == 1 && col == x {
				result += "C"
			} else if row == y && col == 1 {
				result += "C"
			} else if row == y && col == x {
				result += "A"
			} else if row == 1 || row == y || col == 1 || col == x {
				result += "B"
			} else {
				result += " "
			}
		}
		result += "\n"
	}

	return result
}

func getDimensions(input string) (int, int) {
	if input == "" {
		return 0, 0
	}

	width := 0
	height := 0
	currentWidth := 0

	for i := 0; i < len(input); i++ {
		if input[i] == '\n' {
			height++
			if width == 0 {
				width = currentWidth
			}
			currentWidth = 0
		} else {
			currentWidth++
		}
	}

	return width, height
}

func main() {
	data, _ := io.ReadAll(os.Stdin)
	input := string(data)

	x, y := getDimensions(input)

	matches := []string{}

	if input == quadA(x, y) {
		matches = append(matches, fmt.Sprintf("[quadA] [%d] [%d]", x, y))
	}
	if input == quadB(x, y) {
		matches = append(matches, fmt.Sprintf("[quadB] [%d] [%d]", x, y))
	}
	if input == quadC(x, y) {
		matches = append(matches, fmt.Sprintf("[quadC] [%d] [%d]", x, y))
	}
	if input == quadD(x, y) {
		matches = append(matches, fmt.Sprintf("[quadD] [%d] [%d]", x, y))
	}
	if input == quadE(x, y) {
		matches = append(matches, fmt.Sprintf("[quadE] [%d] [%d]", x, y))
	}

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	for i := 0; i < len(matches); i++ {
		if i > 0 {
			fmt.Print(" || ")
		}
		fmt.Print(matches[i])
	}
	fmt.Println()
}
