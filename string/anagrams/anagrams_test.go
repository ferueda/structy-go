package anagrams

import "testing"

var testTable = []struct {
	name     string
	input1   string
	input2   string
	expected bool
}{
	{"test_00", "restful", "fluster", true},
	{"test_01", "cats", "tocs", false},
	{"test_02", "monkeyswrite", "newyorktimes", true},
	{"test_03", "paper", "reapa", false},
	{"test_04", "elbow", "below", true},
	{"test_05", "tax", "taxi", false},
	{"test_06", "taxi", "tax", false},
	{"test_07", "night", "thing", true},
	{"test_08", "abbc", "aabc", false},
}

func TestAnagrams(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := anagrams(test.input1, test.input2)
			if result != test.expected {
				t.Errorf("For inputs '%s' and '%s', expected result is '%v' but got '%v'", test.input1, test.input2, test.expected, result)
			}
		})
	}
}
