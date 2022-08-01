package leafList

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
		[]string{"d", "e", "f"}},

	{"test_01",
		&binarytree.Node[string]{Val: "a",
			Left: &binarytree.Node[string]{Val: "b",
				Left: &binarytree.Node[string]{Val: "d"},
				Right: &binarytree.Node[string]{Val: "e",
					Left: &binarytree.Node[string]{Val: "g"}}},
			Right: &binarytree.Node[string]{Val: "c",
				Right: &binarytree.Node[string]{Val: "f",
					Right: &binarytree.Node[string]{Val: "h"}}}},
		[]string{"d", "g", "h"}},

	{"test_02",
		&binarytree.Node[string]{Val: "5",
			Left: &binarytree.Node[string]{Val: "11",
				Left: &binarytree.Node[string]{Val: "20"},
				Right: &binarytree.Node[string]{Val: "15",
					Left:  &binarytree.Node[string]{Val: "1"},
					Right: &binarytree.Node[string]{Val: "3"}}},
			Right: &binarytree.Node[string]{Val: "54"}},
		[]string{"20", "1", "3", "54"}},

	{"test_03",
		&binarytree.Node[string]{Val: "x"},
		[]string{"x"}},

	{"test_04",
		nil,
		[]string{}},
}

func TestLeafList(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := leafList(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For inputs '%v', expected result is '%v' but got '%v'", test.input, test.expected, result)
			}
		})
	}
}
