package piscine

func SortWordArr(a []string) {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if greater(a[j], a[j+1]) {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func greater(s1, s2 string) bool {
	i := 0

	for i < len(s1) && i < len(s2) {
		if s1[i] > s2[i] {
			return true
		}
		if s1[i] < s2[i] {
			return false
		}
		i++
	}

	return len(s1) > len(s2)
}
