package piscine

func Capitalize(s string) string {
	result := ""
	newWord := true

	for i := 0; i < len(s); i++ {
		c := s[i]

		// Check if alphanumeric
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
			if newWord {
				// First character of word → uppercase if letter
				if c >= 'a' && c <= 'z' {
					c = c - 32
				}
				newWord = false
			} else {
				// Inside word → lowercase if letter
				if c >= 'A' && c <= 'Z' {
					c = c + 32
				}
			}
		} else {
			// Non-alphanumeric → next char starts new word
			newWord = true
		}

		result += string(c)
	}

	return result
}
