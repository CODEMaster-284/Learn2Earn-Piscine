package piscine

func IsAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		c := s[i]

		if !((c >= 'a' && c <= 'z') ||
			(c >= 'A' && c <= 'Z') ||
			(c >= '0' && c <= '9')) {
			return false
		}
	}
	return true
}
