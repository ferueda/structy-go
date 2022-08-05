package islandCount

import (
	"testing"
)

var testTable = []struct {
	name     string
	input    [][]string
	expected int
}{
	{"test_00",
		[][]string{
			{"W", "L", "W", "W", "W"},
			{"W", "L", "W", "W", "W"},
			{"W", "W", "W", "L", "W"},
			{"W", "W", "L", "L", "W"},
			{"L", "W", "W", "L", "L"},
			{"L", "L", "W", "W", "W"},
		},
		3},

	{"test_01",
		[][]string{
			{"L", "W", "W", "L", "W"},
			{"L", "W", "W", "L", "L"},
			{"W", "L", "W", "L", "W"},
			{"W", "W", "W", "W", "W"},
			{"W", "W", "L", "L", "L"},
		},
		4},

	{"test_02",
		[][]string{
			{"L", "L", "L"},
			{"L", "L", "L"},
			{"L", "L", "L"},
		},
		1},

	{"test_03",
		[][]string{
			{"W", "W"},
			{"W", "W"},
			{"W", "W"},
		},
		0},
}

func TestIslandCount(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := islandCount(test.input)
			if result != test.expected {
				t.Errorf("For input '%v', expected result is '%v' but got '%v'", test.input, test.expected, result)
			}
		})
	}
}
