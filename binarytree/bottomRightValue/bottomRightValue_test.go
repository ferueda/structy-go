package bottomRightValue

import (
	"testing"

	"github.com/ferueda/structy-go/binarytree"
)

var testTable = []struct {
	name     string
	input    *binarytree.Node[string]
	expected string
}{
	{"test_00",
		&binarytree.Node[string]{Val: "3",
			Left: &binarytree.Node[string]{Val: "11",
				Left:  &binarytree.Node[string]{Val: "4"},
				Right: &binarytree.Node[string]{Val: "-2"}},
			Right: &binarytree.Node[string]{Val: "10",
				Right: &binarytree.Node[string]{Val: "1"}}},
		"1"},

	{"test_01",
		&binarytree.Node[string]{Val: "-1",
			Left: &binarytree.Node[string]{Val: "-6",
				Left: &binarytree.Node[string]{Val: "-3"},
				Right: &binarytree.Node[string]{Val: "-4",
					Left:  &binarytree.Node[string]{Val: "-2"},
					Right: &binarytree.Node[string]{Val: "6"}}},
			Right: &binarytree.Node[string]{Val: "-5",
				Right: &binarytree.Node[string]{Val: "-13"}}},
		"6"},

	{"test_02",
		&binarytree.Node[string]{Val: "-1",
			Left: &binarytree.Node[string]{Val: "-6",
				Left: &binarytree.Node[string]{Val: "-3"},
				Right: &binarytree.Node[string]{Val: "-4",
					Left:  &binarytree.Node[string]{Val: "-2"},
					Right: &binarytree.Node[string]{Val: "6"}}},
			Right: &binarytree.Node[string]{Val: "-5",
				Right: &binarytree.Node[string]{Val: "-13",
					Left: &binarytree.Node[string]{Val: "7"}}}},
		"7"},

	{"test_03",
		&binarytree.Node[string]{Val: "a",
			Left: &binarytree.Node[string]{Val: "b",
				Right: &binarytree.Node[string]{Val: "d",
					Left: &binarytree.Node[string]{Val: "e",
						Left: &binarytree.Node[string]{Val: "f"}}}},
			Right: &binarytree.Node[string]{Val: "c"}},
		"f"},

	{"test_04",
		&binarytree.Node[string]{Val: "a"},
		"a"},
}

func TestBottomRightValue(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := bottomRightValue(test.input)
			if result != test.expected {
				t.Errorf("For inputs '%v', expected result is '%v' but got '%v'", test.input, test.expected, result)
			}
		})
	}
}
