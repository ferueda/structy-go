package mergeLists

import (
	"reflect"
	"testing"

	"github.com/ferueda/structy-go/linkedlist"
)

var testTable = []struct {
	name     string
	input1   *linkedlist.Node[int]
	input2   *linkedlist.Node[int]
	expected []int
}{
	{"test_00",
		&linkedlist.Node[int]{Val: 5,
			Next: &linkedlist.Node[int]{Val: 7,
				Next: &linkedlist.Node[int]{Val: 10,
					Next: &linkedlist.Node[int]{Val: 12,
						Next: &linkedlist.Node[int]{Val: 20,
							Next: &linkedlist.Node[int]{Val: 28}}}}}},
		&linkedlist.Node[int]{Val: 6,
			Next: &linkedlist.Node[int]{Val: 8,
				Next: &linkedlist.Node[int]{Val: 9,
					Next: &linkedlist.Node[int]{Val: 25}}}},
		[]int{5, 6, 7, 8, 9, 10, 12, 20, 25, 28}},

	{"test_01",
		&linkedlist.Node[int]{Val: 5,
			Next: &linkedlist.Node[int]{Val: 7,
				Next: &linkedlist.Node[int]{Val: 10,
					Next: &linkedlist.Node[int]{Val: 12,
						Next: &linkedlist.Node[int]{Val: 20,
							Next: &linkedlist.Node[int]{Val: 28}}}}}},
		&linkedlist.Node[int]{Val: 1,
			Next: &linkedlist.Node[int]{Val: 8,
				Next: &linkedlist.Node[int]{Val: 9,
					Next: &linkedlist.Node[int]{Val: 10}}}},
		[]int{1, 5, 7, 8, 9, 10, 10, 12, 20, 28}},

	{"test_02",
		&linkedlist.Node[int]{Val: 30},
		&linkedlist.Node[int]{Val: 15,
			Next: &linkedlist.Node[int]{Val: 67}},
		[]int{15, 30, 67}},
}

func TestMergeLists(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := mergeLists(test.input1, test.input2)
			if !reflect.DeepEqual(result.GetValues(), test.expected) {
				t.Errorf("For inputs '%s' and '%s', expected result is '%v' but got '%v'", test.input1.String(), test.input2.String(), test.expected, result.GetValues())
			}
		})
	}

}
