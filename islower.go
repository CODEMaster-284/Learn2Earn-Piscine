package piscine

func IsLower(s string) bool {
	if len(s) == 0 {
		return false
	}

	for i := 0; i < len(s); i++ {
		c := s[i]

		if c < 'a' || c > 'z' {
			return false
		}
	}

	return true
}
