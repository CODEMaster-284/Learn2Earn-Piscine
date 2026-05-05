package piscine

func LoafOfBread(str string) string {
	runes := []rune(str)
	letters := 0

	for _, r := range runes {
		if r != ' ' {
			letters++
		}
	}

	if letters < 5 {
		return "Invalid Output\n"
	}

	result := ""
	first := true

	for i := 0; i < len(runes); {
		word := ""
		count := 0

		for i < len(runes) && count < 5 {
			if runes[i] != ' ' {
				word += string(runes[i])
				count++
			}
			i++
		}

		if word == "" {
			break
		}

		if !first {
			result += " "
		}
		result += word
		first = false

		if count == 5 && i < len(runes) {
			i++
		}
	}

	return result + "\n"
}
