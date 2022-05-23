package zipperLists

import (
	"reflect"
	"testing"

	"github.com/ferueda/structy-go/linkedlist"
)

var testTable = []struct {
	name     string
	input1   *linkedlist.Node[string]
	input2   *linkedlist.Node[string]
	expected []string
}{
	{"test_00",
		&linkedlist.Node[string]{Val: "a",
			Next: &linkedlist.Node[string]{Val: "b",
				Next: &linkedlist.Node[string]{Val: "c"}}},
		&linkedlist.Node[string]{Val: "x",
			Next: &linkedlist.Node[string]{Val: "y",
				Next: &linkedlist.Node[string]{Val: "z"}}},
		[]string{"a", "x", "b", "y", "c", "z"}},

	{"test_01",
		&linkedlist.Node[string]{Val: "a",
			Next: &linkedlist.Node[string]{Val: "b",
				Next: &linkedlist.Node[string]{Val: "c",
					Next: &linkedlist.Node[string]{Val: "d",
						Next: &linkedlist.Node[string]{Val: "e",
							Next: &linkedlist.Node[string]{Val: "f"}}}}}},
		&linkedlist.Node[string]{Val: "x",
			Next: &linkedlist.Node[string]{Val: "y",
				Next: &linkedlist.Node[string]{Val: "z"}}},
		[]string{"a", "x", "b", "y", "c", "z", "d", "e", "f"}},

	{"test_02",
		&linkedlist.Node[string]{Val: "s",
			Next: &linkedlist.Node[string]{Val: "t"}},
		&linkedlist.Node[string]{Val: "1",
			Next: &linkedlist.Node[string]{Val: "2",
				Next: &linkedlist.Node[string]{Val: "3",
					Next: &linkedlist.Node[string]{Val: "4"}}}},
		[]string{"s", "1", "t", "2", "3", "4"}},

	{"test_03",
		&linkedlist.Node[string]{Val: "w"},
		&linkedlist.Node[string]{Val: "1",
			Next: &linkedlist.Node[string]{Val: "2",
				Next: &linkedlist.Node[string]{Val: "3"}}},
		[]string{"w", "1", "2", "3"}},

	{"test_04",
		&linkedlist.Node[string]{Val: "1",
			Next: &linkedlist.Node[string]{Val: "2",
				Next: &linkedlist.Node[string]{Val: "3"}}},
		&linkedlist.Node[string]{Val: "w"},
		[]string{"1", "w", "2", "3"}},
}

func TestZipperList(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := zipperLists(test.input1, test.input2)
			if !reflect.DeepEqual(result.GetValues(), test.expected) {
				t.Errorf("For inputs '%s' and '%s', expected result is '%v' but got '%v'", test.input1.String(), test.input2.String(), test.expected, result.GetValues())
			}
		})
	}

}
