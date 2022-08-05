package undirectedPath

import (
	"testing"
)

var testTable = []struct {
	name     string
	input1   [][]string
	input2   string
	input3   string
	expected bool
}{
	{"test_00",
		[][]string{
			{"i", "j"},
			{"k", "i"},
			{"m", "k"},
			{"k", "l"},
			{"o", "n"},
		},
		"j",
		"m",
		true},

	{"test_01",
		[][]string{
			{"i", "j"},
			{"k", "i"},
			{"m", "k"},
			{"k", "l"},
			{"o", "n"},
		},
		"m",
		"j",
		true},

	{"test_02",
		[][]string{
			{"i", "j"},
			{"k", "i"},
			{"m", "k"},
			{"k", "l"},
			{"o", "n"},
		},
		"l",
		"j",
		true},

	{"test_03",
		[][]string{
			{"i", "j"},
			{"k", "i"},
			{"m", "k"},
			{"k", "l"},
			{"o", "n"},
		},
		"k",
		"o",
		false},

	{"test_04",
		[][]string{
			{"i", "j"},
			{"k", "i"},
			{"m", "k"},
			{"k", "l"},
			{"o", "n"},
		},
		"i",
		"o",
		false},

	{"test_05",
		[][]string{
			{"b", "a"},
			{"c", "a"},
			{"b", "c"},
			{"q", "r"},
			{"q", "s"},
			{"q", "u"},
			{"q", "t"},
		},
		"a",
		"b",
		true},

	{"test_06",
		[][]string{
			{"b", "a"},
			{"c", "a"},
			{"b", "c"},
			{"q", "r"},
			{"q", "s"},
			{"q", "u"},
			{"q", "t"},
		},
		"a",
		"c",
		true},

	{"test_07",
		[][]string{
			{"b", "a"},
			{"c", "a"},
			{"b", "c"},
			{"q", "r"},
			{"q", "s"},
			{"q", "u"},
			{"q", "t"},
		},
		"r",
		"t",
		true},

	{"test_08",
		[][]string{
			{"b", "a"},
			{"c", "a"},
			{"b", "c"},
			{"q", "r"},
			{"q", "s"},
			{"q", "u"},
			{"q", "t"},
		},
		"r",
		"b",
		false},

	{"test_09",
		[][]string{
			{"s", "r"},
			{"t", "q"},
			{"q", "r"},
		},
		"r",
		"t",
		true},
}

func TestUndirectedPath(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := undirectedPath(test.input1, test.input2, test.input3)
			if result != test.expected {
				t.Errorf("For inputs '%v', '%v', and '%v', expected result is '%v' but got '%v'", test.input1, test.input2, test.input3, test.expected, result)
			}
		})
	}
}
