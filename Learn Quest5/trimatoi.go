package piscine

func TrimAtoi(s string) int {
	result := 0
	sign := 1
	foundDigit := false

	for i := 0; i < len(s); i++ {
		c := s[i]

		// Check for sign BEFORE digits
		if c == '-' && !foundDigit {
			sign = -1
		}

		// Check if digit
		if c >= '0' && c <= '9' {
			foundDigit = true
			result = result*10 + int(c-'0')
		}
	}

	return result * sign
}
