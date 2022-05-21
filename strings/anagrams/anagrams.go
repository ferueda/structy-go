package anagrams

func anagrams(s1, s2 string) bool {
	chars := make(map[rune]int)

	for _, c := range s1 {
		chars[c] += 1
	}

	for _, c := range s2 {
		if _, ok := chars[c]; !ok {
			return false
		}

		if chars[c] == 1 {
			delete(chars, c)
		} else {
			chars[c] -= 1
		}
	}

	return len(chars) == 0
}
