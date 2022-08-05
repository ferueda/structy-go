package shortestPath

import (
	"testing"
)

var testTable = []struct {
	name     string
	input1   [][]string
	input2   string
	input3   string
	expected int
}{
	{"test_00",
		[][]string{
			{"w", "x"},
			{"x", "y"},
			{"z", "y"},
			{"z", "v"},
			{"w", "v"},
		},
		"w",
		"z",
		2},

	{"test_01",
		[][]string{
			{"w", "x"},
			{"x", "y"},
			{"z", "y"},
			{"z", "v"},
			{"w", "v"},
		},
		"y",
		"x",
		1},

	{"test_02",
		[][]string{
			{"a", "c"},
			{"a", "b"},
			{"c", "b"},
			{"c", "d"},
			{"b", "d"},
			{"e", "d"},
			{"g", "f"},
		},
		"a",
		"e",
		3},

	{"test_03",
		[][]string{
			{"a", "c"},
			{"a", "b"},
			{"c", "b"},
			{"c", "d"},
			{"b", "d"},
			{"e", "d"},
			{"g", "f"},
		},
		"e",
		"c",
		2},

	{"test_04",
		[][]string{
			{"a", "c"},
			{"a", "b"},
			{"c", "b"},
			{"c", "d"},
			{"b", "d"},
			{"e", "d"},
			{"g", "f"},
		},
		"b",
		"g",
		-1},

	{"test_05",
		[][]string{
			{"c", "n"},
			{"c", "e"},
			{"c", "s"},
			{"c", "w"},
			{"w", "e"},
		},
		"w",
		"e",
		1},

	{"test_06",
		[][]string{
			{"c", "n"},
			{"c", "e"},
			{"c", "s"},
			{"c", "w"},
			{"w", "e"},
		},
		"n",
		"e",
		2},

	{"test_07",
		[][]string{
			{"m", "n"},
			{"n", "o"},
			{"o", "p"},
			{"p", "q"},
			{"t", "o"},
			{"r", "q"},
			{"r", "s"},
		},
		"m",
		"s",
		6},
}

func TestShortestPath(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := shortestPath(test.input1, test.input2, test.input3)
			if result != test.expected {
				t.Errorf("For inputs '%v', '%v', and '%v', expected result is '%v' but got '%v'", test.input1, test.input2, test.input3, test.expected, result)
			}
		})
	}
}
