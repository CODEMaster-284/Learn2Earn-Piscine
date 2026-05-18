package piscine

func StringToIntSlice(str string) []int {
	if str == "" {
		return nil
	}

	runes := []rune(str)
	result := make([]int, len(runes))

	for i, r := range runes {
		result[i] = int(r)
	}

	return result
}
