package fivesort

import (
	"reflect"
	"testing"
)

var testTable = []struct {
	name     string
	input    []int
	expected []int
}{
	{"test_00", []int{12, 5, 1, 5, 12, 7}, []int{12, 7, 1, 12, 5, 5}},
	{"test_01", []int{5, 2, 5, 6, 5, 1, 10, 2, 5, 5}, []int{2, 2, 10, 6, 1, 5, 5, 5, 5, 5}},
	{"test_02", []int{5, 5, 5, 1, 1, 1, 4}, []int{4, 1, 1, 1, 5, 5, 5}},
	{"test_03", []int{5, 5, 6, 5, 5, 5, 5}, []int{6, 5, 5, 5, 5, 5, 5}},
	{"test_04", []int{5, 1, 2, 5, 5, 3, 2, 5, 1, 5, 5, 5, 4, 5}, []int{4, 1, 2, 1, 2, 3, 5, 5, 5, 5, 5, 5, 5, 5}},
}

func TestFiveSort(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := fiveSort(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For input '%v', expected result is '%v' but got '%v'", test.input, test.expected, result)
			}
		})
	}
}
