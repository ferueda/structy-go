package uncompress

import "testing"

var testTable = []struct {
	name     string
	input    string
	expected string
}{
	{"test_00", "2c3a1t", "ccaaat"},
	{"test_01", "4s2b", "ssssbb"},
	{"test_02", "2p1o5p", "ppoppppp"},
	{"test_03", "3n12e2z", "nnneeeeeeeeeeeezz"},
	{"test_04", "127y", "yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy"},
}

func TestUncompress(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := uncompress(test.input)
			if result != test.expected {
				t.Errorf("For input '%s', expected result is '%s' but got '%s'", test.input, test.expected, result)
			}
		})
	}
}
