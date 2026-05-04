package piscine

func Rot14(s string) string {
	r := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			r += string((s[i]-'a'+14)%26 + 'a')
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			r += string((s[i]-'A'+14)%26 + 'A')
		} else {
			r += string(s[i])
		}
	}
	return r
}
