package treesum

import (
	"reflect"
	"testing"

	"github.com/ferueda/structy-go/binarytree"
)

var testTable = []struct {
	name     string
	input    *binarytree.Node[int]
	expected int
}{
	{"test_00",
		&binarytree.Node[int]{Val: 3,
			Left: &binarytree.Node[int]{Val: 11,
				Left:  &binarytree.Node[int]{Val: 4},
				Right: &binarytree.Node[int]{Val: -2}},
			Right: &binarytree.Node[int]{Val: 4,
				Right: &binarytree.Node[int]{Val: 1}}},
		21},

	{"test_01",
		&binarytree.Node[int]{Val: 1,
			Left: &binarytree.Node[int]{Val: 6,
				Left: &binarytree.Node[int]{Val: 3},
				Right: &binarytree.Node[int]{Val: -6,
					Left: &binarytree.Node[int]{Val: 2}}},
			Right: &binarytree.Node[int]{Val: 0,
				Right: &binarytree.Node[int]{Val: 2,
					Right: &binarytree.Node[int]{Val: 2}}}},
		10},

	{"test_02",
		nil,
		0},
}

func TestTreeSum(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := treeSum(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For input '%v', expected result is '%v' but got '%v'", test.input, test.expected, result)
			}
		})
	}
}
