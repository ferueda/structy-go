package removeNode

import (
	"reflect"
	"testing"

	"github.com/ferueda/structy-go/linkedlist"
)

var testTable = []struct {
	name     string
	input1   *linkedlist.Node[string]
	input2   string
	expected []string
}{
	{"test_00",
		&linkedlist.Node[string]{Val: "a",
			Next: &linkedlist.Node[string]{Val: "b",
				Next: &linkedlist.Node[string]{Val: "c",
					Next: &linkedlist.Node[string]{Val: "d",
						Next: &linkedlist.Node[string]{Val: "e",
							Next: &linkedlist.Node[string]{Val: "f"}}}}}},
		"c",
		[]string{"a", "b", "d", "e", "f"}},

	{"test_01",
		&linkedlist.Node[string]{Val: "a",
			Next: &linkedlist.Node[string]{Val: "b",
				Next: &linkedlist.Node[string]{Val: "c"}}},
		"c",
		[]string{"a", "b"}},

	{"test_02",
		&linkedlist.Node[string]{Val: "a",
			Next: &linkedlist.Node[string]{Val: "b",
				Next: &linkedlist.Node[string]{Val: "c"}}},
		"a",
		[]string{"b", "c"}},

	{"test_03",
		&linkedlist.Node[string]{Val: "a",
			Next: &linkedlist.Node[string]{Val: "b",
				Next: &linkedlist.Node[string]{Val: "c",
					Next: &linkedlist.Node[string]{Val: "b"}}}},
		"b",
		[]string{"a", "c", "b"}},

	{"test_04",
		&linkedlist.Node[string]{Val: "a"},
		"a",
		[]string{}},
}

func TestRemoveNode(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := removeNode(test.input1, test.input2)
			if !reflect.DeepEqual(result.GetValues(), test.expected) {
				t.Errorf("For inputs '%s' and '%s', expected result is '%v' but got '%v'", test.input1.String(), test.input2, test.expected, result.GetValues())
			}
		})
	}

}
