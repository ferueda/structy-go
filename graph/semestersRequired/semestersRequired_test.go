package semestersRequired

import (
	"testing"
)

var testTable = []struct {
	name     string
	input    [][]int
	expected int
}{
	{"test_00",
		[][]int{
			{1, 2},
			{2, 4},
			{3, 5},
			{0, 5},
		},
		3},

	{"test_01",
		[][]int{
			{4, 3},
			{3, 2},
			{2, 1},
			{1, 0},
			{5, 2},
			{5, 6},
		},
		5},

	{"test_02",
		[][]int{
			{1, 0},
			{3, 4},
			{1, 2},
			{3, 2},
		},
		2},

	{"test_03",
		[][]int{},
		1},

	{"test_04",
		[][]int{
			{0, 2},
			{0, 1},
			{1, 2},
		},
		3},

	{"test_05",
		[][]int{
			{3, 4},
			{3, 0},
			{3, 1},
			{3, 2},
			{3, 5},
		},
		2},
}

func TestSemestersRequired(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := semestersRequired(test.input)
			if result != test.expected {
				t.Errorf("For input '%v', expected result is '%v' but got '%v'", test.input, test.expected, result)
			}
		})
	}
}
