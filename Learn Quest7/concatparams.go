package piscine

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}

	// First calculate total length
	totalLen := 0
	for i := 0; i < len(args); i++ {
		totalLen += len(args[i])
	}
	totalLen += len(args) - 1 // for '\n'

	// Create byte slice
	result := make([]byte, totalLen)

	index := 0
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(args[i]); j++ {
			result[index] = args[i][j]
			index++
		}
		if i != len(args)-1 {
			result[index] = '\n'
			index++
		}
	}

	return string(result)
}
