package piscine

func ToUpper(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		c := s[i]

		if c >= 'a' && c <= 'z' {
			c = c - 32
		}

		result += string(c)
	}

	return result
}
