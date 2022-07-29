package allTreePaths

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
		[][]string{{"a", "b", "d"}, {"a", "b", "e"}, {"a", "c", "f"}}},

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
		[][]string{{"a", "b", "d"}, {"a", "b", "e", "g"}, {"a", "b", "e", "h"}, {"a", "c", "f", "i"}}},

	{"test_02",
		&binarytree.Node[string]{Val: "q",
			Left: &binarytree.Node[string]{Val: "r",
				Right: &binarytree.Node[string]{Val: "t",
					Left: &binarytree.Node[string]{Val: "u",
						Left: &binarytree.Node[string]{Val: "v"}}}},
			Right: &binarytree.Node[string]{Val: "s"}},
		[][]string{{"q", "r", "t", "u", "v"}, {"q", "s"}}},

	{"test_03",
		&binarytree.Node[string]{Val: "z"},
		[][]string{{"z"}}},
}

func TestAllTreePaths(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := allTreePaths(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For inputs '%v', expected result is '%v' but got '%v'", test.input, test.expected, result)
			}
		})
	}
}
