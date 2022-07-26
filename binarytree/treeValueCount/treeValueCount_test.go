package treeValueCount

import (
	"reflect"
	"testing"

	"github.com/ferueda/structy-go/binarytree"
)

var testTable = []struct {
	name     string
	input1   *binarytree.Node[int]
	input2   int
	expected int
}{
	{"test_00",
		&binarytree.Node[int]{Val: 12,
			Left: &binarytree.Node[int]{Val: 6,
				Left:  &binarytree.Node[int]{Val: 4},
				Right: &binarytree.Node[int]{Val: 6}},
			Right: &binarytree.Node[int]{Val: 6,
				Right: &binarytree.Node[int]{Val: 12}}},
		6,
		3},

	{"test_01",
		&binarytree.Node[int]{Val: 12,
			Left: &binarytree.Node[int]{Val: 6,
				Left:  &binarytree.Node[int]{Val: 4},
				Right: &binarytree.Node[int]{Val: 6}},
			Right: &binarytree.Node[int]{Val: 6,
				Right: &binarytree.Node[int]{Val: 12}}},
		12,
		2},

	{"test_02",
		&binarytree.Node[int]{Val: 7,
			Left: &binarytree.Node[int]{Val: 5,
				Left: &binarytree.Node[int]{Val: 1},
				Right: &binarytree.Node[int]{Val: 8,
					Left: &binarytree.Node[int]{Val: 1}}},
			Right: &binarytree.Node[int]{Val: 1,
				Right: &binarytree.Node[int]{Val: 7,
					Right: &binarytree.Node[int]{Val: 1}}}},
		1,
		4},

	{"test_03",
		&binarytree.Node[int]{Val: 7,
			Left: &binarytree.Node[int]{Val: 5,
				Left: &binarytree.Node[int]{Val: 1},
				Right: &binarytree.Node[int]{Val: 8,
					Left: &binarytree.Node[int]{Val: 1}}},
			Right: &binarytree.Node[int]{Val: 1,
				Right: &binarytree.Node[int]{Val: 7,
					Right: &binarytree.Node[int]{Val: 1}}}},
		9,
		0},

	{"test_04",
		nil,
		42,
		0},
}

func TestTreeValueCount(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := treeValueCount(test.input1, test.input2)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For inputs '%v' and '%v', expected result is '%v' but got '%v'", test.input1, test.input2, test.expected, result)
			}
		})
	}
}
