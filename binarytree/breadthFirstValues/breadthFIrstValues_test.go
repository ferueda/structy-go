package breadthFirstValues

import (
	"reflect"
	"testing"

	"github.com/ferueda/structy-go/binarytree"
)

var testTable = []struct {
	name     string
	input    *binarytree.Node[string]
	expected []string
}{
	{"test_00",
		&binarytree.Node[string]{Val: "a",
			Left: &binarytree.Node[string]{Val: "b",
				Left:  &binarytree.Node[string]{Val: "d"},
				Right: &binarytree.Node[string]{Val: "e"}},
			Right: &binarytree.Node[string]{Val: "c",
				Right: &binarytree.Node[string]{Val: "f"}}},
		[]string{"a", "b", "c", "d", "e", "f"}},

	{"test_01",
		&binarytree.Node[string]{Val: "a",
			Left: &binarytree.Node[string]{Val: "b",
				Left: &binarytree.Node[string]{Val: "d"},
				Right: &binarytree.Node[string]{Val: "e",
					Left: &binarytree.Node[string]{Val: "g"}}},
			Right: &binarytree.Node[string]{Val: "c",
				Right: &binarytree.Node[string]{Val: "f",
					Right: &binarytree.Node[string]{Val: "h"}}}},
		[]string{"a", "b", "c", "d", "e", "f", "g", "h"}},

	{"test_02",
		&binarytree.Node[string]{Val: "a"},
		[]string{"a"}},

	{"test_03",
		&binarytree.Node[string]{Val: "a",
			Right: &binarytree.Node[string]{Val: "b",
				Left: &binarytree.Node[string]{Val: "c",
					Left: &binarytree.Node[string]{Val: "x"},
					Right: &binarytree.Node[string]{Val: "d",
						Right: &binarytree.Node[string]{Val: "e"}}}}},
		[]string{"a", "b", "c", "x", "d", "e"}},

	{"test_04",
		nil,
		[]string{}},
}

func TestBreadthFirstValues(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := breadthFirstValues(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For input '%v', expected result is '%v' but got '%v'", test.input, test.expected, result)
			}
		})
	}
}
