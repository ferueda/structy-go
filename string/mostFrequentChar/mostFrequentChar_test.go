package mostFrequentChar

import "testing"

var testTable = []struct {
	name     string
	input    string
	expected string
}{
	{"test_00", "bookeeper", "e"},
	{"test_01", "david", "d"},
	{"test_02", "abby", "b"},
	{"test_03", "mississippi", "i"},
	{"test_04", "potato", "o"},
	{"test_05", "eleventennine", "e"},
	{"test_06", "riverbed", "r"},
}

func TestMostFrequentChar(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := mostFrequentChar(test.input)
			if result != test.expected {
				t.Errorf("For input '%s', expected result is '%v' but got '%v'", test.input, test.expected, result)
			}
		})
	}
}
