package uncompress

import (
	"strconv"
	"strings"
	"unicode"
)

func uncompress(s string) string {
	var r []string
	i, j := 0, 0
	sr := []rune(s)

	for j < len(s) {
		if unicode.IsDigit(sr[j]) {
			j += 1
		} else {
			num, _ := strconv.Atoi(string(sr[i:j]))
			r = append(r, strings.Repeat(string(sr[j]), num))
			j += 1
			i = j
		}
	}

	return strings.Join(r, "")
}
