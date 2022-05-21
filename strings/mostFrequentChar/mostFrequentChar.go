package mostFrequentChar

func mostFrequentChar(s string) string {
	chars := make(map[rune]int)

	for _, c := range s {
		chars[c] += 1
	}

	var r rune

	for _, c := range s {
		if chars[c] > chars[r] {
			r = c
		}
	}

	return string(r)
}
