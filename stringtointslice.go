package piscine

func StringToIntSlice(str string) []int {
	result := make([]int, len(str))

	for i := 0; i < len(str); i++ {
		result[i] = int(str[i])
	}

	return result
}
