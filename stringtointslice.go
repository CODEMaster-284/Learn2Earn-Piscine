package piscine

func StringToIntSlice(str string) []int {
	runes := []rune(str)
	result := make([]int, len(runes))

	for i, r := range runes {
		result[i] = int(r)
	}

	return result
}
