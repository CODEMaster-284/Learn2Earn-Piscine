package piscine

func AlphaCount(s string) int {
	count := 0

	for i := 0; i < len(s); i++ {
		c := s[i]

		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			count++
		}
	}

	return count
}
