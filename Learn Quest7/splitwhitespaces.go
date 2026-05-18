package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	start := -1

	for i := 0; i < len(s); i++ {
		// Check if current char is NOT a whitespace
		if s[i] != ' ' && s[i] != '\t' && s[i] != '\n' {
			if start == -1 {
				start = i
			}
		} else {
			// We hit a separator
			if start != -1 {
				result = append(result, s[start:i])
				start = -1
			}
		}
	}

	// Add last word if string didn't end with space
	if start != -1 {
		result = append(result, s[start:])
	}

	return result
}
