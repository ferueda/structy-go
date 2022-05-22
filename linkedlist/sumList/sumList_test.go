package sumList

import (
	"reflect"
	"testing"

	"github.com/ferueda/structy-go/linkedlist"
)

var testTable = []struct {
	name     string
	input    *linkedlist.Node[int]
	expected int
}{
	{"test_00",
		&linkedlist.Node[int]{Val: 2,
			Next: &linkedlist.Node[int]{Val: 8,
				Next: &linkedlist.Node[int]{Val: 3,
					Next: &linkedlist.Node[int]{Val: -1,
						Next: &linkedlist.Node[int]{Val: 7}}}}},
		19},

	{"test_01",
		&linkedlist.Node[int]{Val: 38,
			Next: &linkedlist.Node[int]{Val: 4}},
		42},

	{"test_02",
		&linkedlist.Node[int]{Val: 100},
		100},

	{"test_03",
		nil,
		0},
}

func TestSumList(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := sumList(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For input '%s', expected result is '%v' but got '%v'", test.input.String(), test.expected, result)
			}
		})
	}
}
