package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb%2 == 0 {
		return false
	}

	i := 3
	for i*i <= nb {
		if nb%i == 0 {
			return false
		}
		i += 2
	}
	return true
}
