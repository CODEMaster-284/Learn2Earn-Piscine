package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	word := ""

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			if word != "" {
				summary[word]++
				word = ""
			}
		} else {
			word += string(str[i])
		}
	}

	if word != "" {
		summary[word]++
	}

	return summary
}
