package piscine

func JumpOver(str string) string {
	runes := []rune(str)

	if len(runes) < 3 {
		return "\n"
	}

	result := ""
	for i := 2; i < len(runes); i += 3 {
		result += string(runes[i])
	}

	return result + "\n"
}
