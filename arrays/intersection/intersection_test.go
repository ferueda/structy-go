package intersection

import (
	"reflect"
	"testing"
)

var testTable = []struct {
	name     string
	input1   []int
	input2   []int
	expected []int
}{
	{"test_00", []int{4, 2, 1, 6}, []int{3, 6, 9, 2, 10}, []int{2, 6}},
	{"test_01", []int{2, 4, 6}, []int{4, 2}, []int{2, 4}},
	{"test_02", []int{4, 2, 1}, []int{1, 2, 4, 6}, []int{4, 2, 1}},
	{"test_03", []int{0, 1, 2}, []int{10, 11}, []int{}},
	{"test_04", generateSlice(), generateSlice(), generateSlice()},
}

func generateSlice() []int {
	a := make([]int, 0, 50000)

	for i := 0; i < 50000; i++ {
		a = append(a, i)
	}

	return a
}

func TestIntersection(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := intersection(test.input1, test.input2)
			if !reflect.DeepEqual(result, test.expected) && len(result) == 0 && len(test.expected) != 0 {
				t.Errorf("For inputs '%v' and '%v', expected result is '%v' but got '%v'", test.input1, test.input2, test.expected, result)
			}
		})
	}
}
