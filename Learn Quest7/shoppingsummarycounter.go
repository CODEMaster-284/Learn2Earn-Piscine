package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)

	for _, item := range Split(str, " ") {
		summary[item]++
	}

	return summary
}
