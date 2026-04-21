package piscine

func UltimateDivMod(a *int, b *int) {
	i := *a
	j := *b
	*a = i / j
	*b = i % j
}
