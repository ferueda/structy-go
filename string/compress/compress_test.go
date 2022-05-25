package compress

import "testing"

var testTable = []struct {
	name     string
	input    string
	expected string
}{
	{"test_00", "ccaaatsss", "2c3at3s"},
	{"test_01", "ssssbbz", "4s2bz"},
	{"test_02", "ppoppppp", "2po5p"},
	{"test_03", "nnneeeeeeeeeeeezz", "3n12e2z"},
	{"test_04", "yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy", "127y"},
}

func TestCompress(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := compress(test.input)
			if result != test.expected {
				t.Errorf("For input '%s', expected result is '%s' but got '%s'", test.input, test.expected, result)
			}
		})
	}
}
