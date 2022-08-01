package treeLevels

import (
	"reflect"
	"testing"

	"github.com/ferueda/structy-go/binarytree"
)

var testTable = []struct {
	name     string
	input    *binarytree.Node[string]
	expected [][]string
}{
	{"test_00",
		&binarytree.Node[string]{Val: "a",
			Left: &binarytree.Node[string]{Val: "b",
				Left:  &binarytree.Node[string]{Val: "d"},
				Right: &binarytree.Node[string]{Val: "e"}},
			Right: &binarytree.Node[string]{Val: "c",
				Right: &binarytree.Node[string]{Val: "f"}}},
		[][]string{{"a"}, {"b", "c"}, {"d", "e", "f"}}},

	{"test_01",
		&binarytree.Node[string]{Val: "a",
			Left: &binarytree.Node[string]{Val: "b",
				Left: &binarytree.Node[string]{Val: "d"},
				Right: &binarytree.Node[string]{Val: "e",
					Left:  &binarytree.Node[string]{Val: "g"},
					Right: &binarytree.Node[string]{Val: "h"}}},
			Right: &binarytree.Node[string]{Val: "c",
				Right: &binarytree.Node[string]{Val: "f",
					Left: &binarytree.Node[string]{Val: "i"}}}},
		[][]string{{"a"}, {"b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}}},

	{"test_02",
		&binarytree.Node[string]{Val: "q",
			Left: &binarytree.Node[string]{Val: "r",
				Right: &binarytree.Node[string]{Val: "t",
					Left: &binarytree.Node[string]{Val: "u",
						Left: &binarytree.Node[string]{Val: "v"}}}},
			Right: &binarytree.Node[string]{Val: "s"}},
		[][]string{{"q"}, {"r", "s"}, {"t"}, {"u"}, {"v"}}},

	{"test_03",
		nil,
		[][]string{}},
}

func TestTreeLevels(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := treeLevels(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For inputs '%v', expected result is '%v' but got '%v'", test.input, test.expected, result)
			}
		})
	}
}
