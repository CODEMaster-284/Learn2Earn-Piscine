package piscine

func Join(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}

	total := len(sep) * (len(strs) - 1)
	for _, str := range strs {
		total += len(str)
	}

	result := make([]byte, 0, total)
	for i, str := range strs {
		if i > 0 {
			result = append(result, sep...)
		}
		result = append(result, str...)
	}

	return string(result)
}
