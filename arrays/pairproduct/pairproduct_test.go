package pairproduct

import "testing"

var testTable = []struct {
	name     string
	input1   []int
	input2   int
	expected [2]int
}{
	{"test_00", []int{3, 2, 5, 4, 1}, 8, [2]int{1, 3}},
	{"test_01", []int{3, 2, 5, 4, 1}, 10, [2]int{1, 2}},
	{"test_01", []int{4, 7, 9, 2, 5, 1}, 5, [2]int{4, 5}},
	{"test_02", []int{4, 7, 9, 2, 5, 1}, 35, [2]int{1, 4}},
	{"test_02", []int{3, 2, 5, 4, 1}, 10, [2]int{1, 2}},
	{"test_02", []int{4, 6, 8, 2}, 16, [2]int{2, 3}},
}

func TestCompress(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := pairProduct(test.input1, test.input2)
			if result != test.expected {
				t.Errorf("For inputs '%v' and '%v', expected result is '%v' but got '%v'", test.input1, test.input2, test.expected, result)
			}
		})
	}
}
