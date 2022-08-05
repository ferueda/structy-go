package closestCarrot

import (
	"testing"
)

var testTable = []struct {
	name     string
	input1   [][]string
	input2   int
	input3   int
	expected int
}{
	{"test_00",
		[][]string{
			{"O", "O", "O", "O", "O"},
			{"O", "X", "O", "O", "O"},
			{"O", "X", "X", "O", "O"},
			{"O", "X", "C", "O", "O"},
			{"O", "X", "X", "O", "O"},
			{"C", "O", "O", "O", "O"},
		},
		1,
		2,
		4},

	{"test_01",
		[][]string{
			{"O", "O", "O", "O", "O"},
			{"O", "X", "O", "O", "O"},
			{"O", "X", "X", "O", "O"},
			{"O", "X", "C", "O", "O"},
			{"O", "X", "X", "O", "O"},
			{"C", "O", "O", "O", "O"},
		},
		0,
		0,
		5},

	{"test_02",
		[][]string{
			{"O", "O", "X", "X", "X"},
			{"O", "X", "X", "X", "C"},
			{"O", "X", "O", "X", "X"},
			{"O", "O", "O", "O", "O"},
			{"O", "X", "X", "X", "X"},
			{"O", "O", "O", "O", "O"},
			{"O", "O", "C", "O", "O"},
			{"O", "O", "O", "O", "O"},
		},
		3,
		4,
		9},

	{"test_03",
		[][]string{
			{"O", "O", "X", "O", "O"},
			{"O", "X", "X", "X", "O"},
			{"O", "X", "C", "C", "O"},
		},
		1,
		4,
		2},

	{"test_04",
		[][]string{
			{"O", "O", "X", "O", "O"},
			{"O", "X", "X", "X", "O"},
			{"O", "X", "C", "C", "O"},
		},
		2,
		0,
		-1},

	{"test_05",
		[][]string{
			{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "X", "X"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "X", "C"},
		},
		0,
		0,
		-1},
}

func TestClosestCarrot(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := closestCarrot(test.input1, test.input2, test.input3)
			if result != test.expected {
				t.Errorf("For inputs '%v', '%v', and '%v', expected result is '%v' but got '%v'", test.input1, test.input2, test.input3, test.expected, result)
			}
		})
	}
}
