package treeMinValue

import (
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
		-2},

	{"test_01",
		&binarytree.Node[int]{Val: 5,
			Left: &binarytree.Node[int]{Val: 11,
				Left:  &binarytree.Node[int]{Val: 4},
				Right: &binarytree.Node[int]{Val: 14}},
			Right: &binarytree.Node[int]{Val: 3,
				Right: &binarytree.Node[int]{Val: 12}}},
		3},

	{"test_02",
		&binarytree.Node[int]{Val: -1,
			Left: &binarytree.Node[int]{Val: -6,
				Left: &binarytree.Node[int]{Val: -3},
				Right: &binarytree.Node[int]{Val: -4,
					Left: &binarytree.Node[int]{Val: -2}}},
			Right: &binarytree.Node[int]{Val: -5,
				Right: &binarytree.Node[int]{Val: -13,
					Right: &binarytree.Node[int]{Val: -2}}}},
		-13},

	{"test_03",
		&binarytree.Node[int]{Val: 42},
		42,
	},
}

func TestTreeMinValue(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := treeMinValue(test.input)
			if result != test.expected {
				t.Errorf("For inputs '%v', expected result is '%v' but got '%v'", test.input, test.expected, result)
			}
		})
	}
}
