package piscine

func Sqrt(nb int) int {
	if nb <= 0 {
		return 0
	}

	i := 1
	for i*i <= nb {
		if i*i == nb {
			return i
		}
		i++
	}
	return 0
}
