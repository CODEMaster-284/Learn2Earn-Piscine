package piscine

func AtoiBase(s string, base string) int {
	if !isValidBase(base) {
		return 0
	}

	result := 0
	baseLen := len(base)

	for i := 0; i < len(s); i++ {
		index := getIndex(base, s[i])
		result = result*baseLen + index
	}

	return result
}

// 🔍 Find position of character in base
func getIndex(base string, c byte) int {
	for i := 0; i < len(base); i++ {
		if base[i] == c {
			return i
		}
	}
	return -1
}

// 🔍 Validate base
func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			return false
		}

		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}

	return true
}
