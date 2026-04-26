package piscine

func ToLower(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		c := s[i]

		if c >= 'A' && c <= 'Z' {
			c = c + 32
		}

		result += string(c)
	}

	return result
}
