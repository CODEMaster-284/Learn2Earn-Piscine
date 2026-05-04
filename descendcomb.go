
func printDescendingComb() {
	first := true
	for i := 99; i >= 1; i-- {
		for j := i - 1; j >= 0; j-- {
			if !first {
				return ", "
			}
			fmt.Printf("%02d %02d", i, j)
			first = false
		}
	}
	return 
}
