package maxpathsum

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
		18},

	{"test_01",
		&binarytree.Node[int]{Val: 5,
			Left: &binarytree.Node[int]{Val: 11,
				Left: &binarytree.Node[int]{Val: 20},
				Right: &binarytree.Node[int]{Val: 15,
					Left:  &binarytree.Node[int]{Val: 1},
					Right: &binarytree.Node[int]{Val: 3}}},
			Right: &binarytree.Node[int]{Val: 54}},
		59},

	{"test_02",
		&binarytree.Node[int]{Val: -1,
			Left: &binarytree.Node[int]{Val: -6,
				Left: &binarytree.Node[int]{Val: -3},
				Right: &binarytree.Node[int]{Val: 0,
					Left: &binarytree.Node[int]{Val: -1}}},
			Right: &binarytree.Node[int]{Val: -5,
				Right: &binarytree.Node[int]{Val: -13,
					Right: &binarytree.Node[int]{Val: -2}}}},
		-8},

	{"test_03",
		&binarytree.Node[int]{Val: 42},
		42,
	},
}

func TestMaxPathSum(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := maxPathSum(test.input)
			if result != test.expected {
				t.Errorf("For inputs '%v', expected result is '%v' but got '%v'", test.input, test.expected, result)
			}
		})
	}
}
