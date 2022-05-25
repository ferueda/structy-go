package compress

import (
	"strconv"
	"strings"
)

func compress(s string) string {
	s += "!"
	var r []string
	i, j := 0, 0

	for j < len(s) {
		if s[j] == s[i] {
			j += 1
		} else {
			if j-i == 1 {
				r = append(r, string(s[i]))
			} else {
				r = append(r, strconv.Itoa(j-i), string(s[i]))
			}

			i = j
		}
	}

	return strings.Join(r, "")
}
