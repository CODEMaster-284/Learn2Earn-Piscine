package piscine

func ActiveBits(n int) int {
	count := 0
	value := uint(n)

	for value > 0 {
		count += int(value & 1)
		value >>= 1
	}

	return count
}
