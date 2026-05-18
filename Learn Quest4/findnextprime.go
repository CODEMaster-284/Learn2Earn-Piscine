package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}

	for {
		isPrime := true

		if nb%2 == 0 && nb != 2 {
			isPrime = false
		} else {
			i := 3
			for i*i <= nb {
				if nb%i == 0 {
					isPrime = false
					break
				}
				i += 2
			}
		}

		if isPrime {
			return nb
		}
		nb++
	}
}
