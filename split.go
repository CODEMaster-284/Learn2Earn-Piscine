package piscine

func Split(s, sep string) []string {
	if sep == "" {
		return nil
	}

	var result []string
	start := 0
	i := 0

	for i <= len(s)-len(sep) {
		if s[i:i+len(sep)] == sep {
			result = append(result, s[start:i])
			i += len(sep)
			start = i
		} else {
			i++
		}
	}

	// Add the last part
	result = append(result, s[start:])

	return result
}
