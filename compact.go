package piscine

func Compact(ptr *[]string) int {
	slice := *ptr
	count := 0

	for i := 0; i < len(slice); i++ {
		if slice[i] != "" {
			slice[count] = slice[i]
			count++
		}
	}

	*ptr = slice[:count]

	return count
}
