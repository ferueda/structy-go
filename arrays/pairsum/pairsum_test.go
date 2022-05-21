package pairsum

import "testing"

var testTable = []struct {
	name     string
	input1   []int
	input2   int
	expected [2]int
}{
	{"test_00", []int{3, 2, 5, 4, 1}, 8, [2]int{0, 2}},
	{"test_01", []int{4, 7, 9, 2, 5, 1}, 5, [2]int{0, 5}},
	{"test_02", []int{4, 7, 9, 2, 5, 1}, 3, [2]int{3, 5}},
	{"test_03", []int{1, 6, 7, 2}, 13, [2]int{1, 2}},
	{"test_04", []int{9, 9}, 18, [2]int{0, 1}},
	{"test_05", []int{6, 4, 2, 8}, 12, [2]int{1, 3}},
}

func TestPairSum(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := pairSum(test.input1, test.input2)
			if result != test.expected {
				t.Errorf("For inputs '%v' and '%v', expected result is '%v' but got '%v'", test.input1, test.input2, test.expected, result)
			}
		})
	}
}
