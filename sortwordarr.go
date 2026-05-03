package piscine

func SortWordArr(a []string) {
	n := len(a)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if compare(a[j], a[j+1]) > 0 {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func compare(s1, s2 string) int {
	i := 0

	for i < len(s1) && i < len(s2) {
		if s1[i] != s2[i] {
			return int(s1[i] - s2[i])
		}
		i++
	}

	return len(s1) - len(s2)
}
