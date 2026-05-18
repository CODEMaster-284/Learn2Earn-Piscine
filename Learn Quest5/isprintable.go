package piscine

func IsPrintable(s string) bool {
	for i := 0; i < len(s); i++ {
		c := s[i]

		if c < 32 || c > 126 {
			return false
		}
	}

	return true
}
