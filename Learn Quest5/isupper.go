package piscine

func IsUpper(s string) bool {
	if len(s) == 0 {
		return false
	}

	for i := 0; i < len(s); i++ {
		c := s[i]

		if c < 'A' || c > 'Z' {
			return false
		}
	}

	return true
}
