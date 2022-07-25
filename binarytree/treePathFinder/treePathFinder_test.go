package treePathFinder

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/ferueda/structy-go/binarytree"
)

var testTable = []struct {
	name     string
	input1   *binarytree.Node[string]
	input2   string
	expected []string
}{
	{"test_00",
		&binarytree.Node[string]{Val: "a",
			Left: &binarytree.Node[string]{Val: "b",
				Left:  &binarytree.Node[string]{Val: "d"},
				Right: &binarytree.Node[string]{Val: "e"}},
			Right: &binarytree.Node[string]{Val: "c",
				Right: &binarytree.Node[string]{Val: "f"}}},
		"e",
		[]string{"a", "b", "e"}},

	{"test_01",
		&binarytree.Node[string]{Val: "a",
			Left: &binarytree.Node[string]{Val: "b",
				Left:  &binarytree.Node[string]{Val: "d"},
				Right: &binarytree.Node[string]{Val: "e"}},
			Right: &binarytree.Node[string]{Val: "c",
				Right: &binarytree.Node[string]{Val: "f"}}},
		"p",
		nil},

	{"test_02",
		&binarytree.Node[string]{Val: "a",
			Left: &binarytree.Node[string]{Val: "b",
				Left: &binarytree.Node[string]{Val: "d"},
				Right: &binarytree.Node[string]{Val: "e",
					Left: &binarytree.Node[string]{Val: "g"}}},
			Right: &binarytree.Node[string]{Val: "c",
				Right: &binarytree.Node[string]{Val: "f",
					Right: &binarytree.Node[string]{Val: "h"}}}},
		"c",
		[]string{"a", "c"}},

	{"test_03",
		&binarytree.Node[string]{Val: "a",
			Left: &binarytree.Node[string]{Val: "b",
				Left: &binarytree.Node[string]{Val: "d"},
				Right: &binarytree.Node[string]{Val: "e",
					Left: &binarytree.Node[string]{Val: "g"}}},
			Right: &binarytree.Node[string]{Val: "c",
				Right: &binarytree.Node[string]{Val: "f",
					Right: &binarytree.Node[string]{Val: "h"}}}},
		"h",
		[]string{"a", "c", "f", "h"}},

	{"test_04",
		&binarytree.Node[string]{Val: "x"},
		"x",
		[]string{"x"}},

	{"test_05",
		nil,
		"x",
		nil},

	{"test_06",
		func() *binarytree.Node[string] {
			root := &binarytree.Node[string]{Val: "0"}
			curr := root
			for i := 1; i <= 6000; i++ {
				curr.Right = &binarytree.Node[string]{Val: strconv.Itoa(i)}
				curr = curr.Right
			}
			return root
		}(),
		"3451",
		func() []string {
			r := make([]string, 0, 3452)
			for i := 0; i <= 3451; i++ {
				r = append(r, strconv.Itoa(i))
			}
			return r
		}()},
}

func TestPathFinder(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := pathFinder(test.input1, test.input2)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For inputs '%v' and '%s', expected result is '%v' but got '%v'", test.input1, test.input2, test.expected, result)
			}
		})
	}
}
