package addLists

import (
	"reflect"
	"testing"

	"github.com/ferueda/structy-go/linkedlist"
)

var testTable = []struct {
	name     string
	input1   *linkedlist.Node[int]
	input2   *linkedlist.Node[int]
	expected *linkedlist.Node[int]
}{
	{"test_00",
		&linkedlist.Node[int]{Val: 1,
			Next: &linkedlist.Node[int]{Val: 2,
				Next: &linkedlist.Node[int]{Val: 6}}},
		&linkedlist.Node[int]{Val: 4,
			Next: &linkedlist.Node[int]{Val: 5,
				Next: &linkedlist.Node[int]{Val: 3}}},
		&linkedlist.Node[int]{Val: 5,
			Next: &linkedlist.Node[int]{Val: 7,
				Next: &linkedlist.Node[int]{Val: 9}}}},

	{"test_01",
		&linkedlist.Node[int]{Val: 1,
			Next: &linkedlist.Node[int]{Val: 4,
				Next: &linkedlist.Node[int]{Val: 5,
					Next: &linkedlist.Node[int]{Val: 7}}}},
		&linkedlist.Node[int]{Val: 2,
			Next: &linkedlist.Node[int]{Val: 3}},
		&linkedlist.Node[int]{Val: 3,
			Next: &linkedlist.Node[int]{Val: 7,
				Next: &linkedlist.Node[int]{Val: 5,
					Next: &linkedlist.Node[int]{Val: 7}}}}},

	{"test_02",
		&linkedlist.Node[int]{Val: 9,
			Next: &linkedlist.Node[int]{Val: 3}},
		&linkedlist.Node[int]{Val: 7,
			Next: &linkedlist.Node[int]{Val: 4}},
		&linkedlist.Node[int]{Val: 6,
			Next: &linkedlist.Node[int]{Val: 8}}},

	{"test_03",
		&linkedlist.Node[int]{Val: 9,
			Next: &linkedlist.Node[int]{Val: 8}},
		&linkedlist.Node[int]{Val: 7,
			Next: &linkedlist.Node[int]{Val: 4}},
		&linkedlist.Node[int]{Val: 6,
			Next: &linkedlist.Node[int]{Val: 3,
				Next: &linkedlist.Node[int]{Val: 1}}}},

	{"test_04",
		&linkedlist.Node[int]{Val: 9,
			Next: &linkedlist.Node[int]{Val: 9,
				Next: &linkedlist.Node[int]{Val: 9}}},
		&linkedlist.Node[int]{Val: 6},
		&linkedlist.Node[int]{Val: 5,
			Next: &linkedlist.Node[int]{Val: 0,
				Next: &linkedlist.Node[int]{Val: 0,
					Next: &linkedlist.Node[int]{Val: 1}}}}},
}

func TestAddLists(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := addLists(test.input1, test.input2)
			if !reflect.DeepEqual(result.GetValues(), test.expected.GetValues()) {
				t.Errorf("For inputs '%s' and '%s', expected result is '%s' but got '%s'", test.input1.String(), test.input2.String(), test.expected.String(), result.String())
			}
		})
	}

}
