package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	decimal := atoiBase(nbr, baseFrom)
	return itoaBase(decimal, baseTo)
}

// Convert string in baseFrom → int
func atoiBase(s, base string) int {
	result := 0
	baseLen := len(base)

	for i := 0; i < len(s); i++ {
		value := indexOf(s[i], base)
		result = result*baseLen + value
	}
	return result
}

// Convert int → string in baseTo
func itoaBase(n int, base string) string {
	if n == 0 {
		return string(base[0])
	}

	baseLen := len(base)
	var result []byte

	for n > 0 {
		remainder := n % baseLen
		result = append([]byte{base[remainder]}, result...)
		n /= baseLen
	}

	return string(result)
}

// Find index of a character in base
func indexOf(c byte, base string) int {
	for i := 0; i < len(base); i++ {
		if base[i] == c {
			return i
		}
	}
	return -1
}
