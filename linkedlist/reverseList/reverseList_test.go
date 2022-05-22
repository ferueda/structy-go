package reverseList

import (
	"reflect"
	"testing"

	"github.com/ferueda/structy-go/linkedlist"
)

var testTable = []struct {
	name     string
	input    *linkedlist.Node[string]
	expected []string
}{
	{"test_00",
		&linkedlist.Node[string]{Val: "a",
			Next: &linkedlist.Node[string]{Val: "b",
				Next: &linkedlist.Node[string]{Val: "c",
					Next: &linkedlist.Node[string]{Val: "d",
						Next: &linkedlist.Node[string]{Val: "e",
							Next: &linkedlist.Node[string]{Val: "f"}}}}}},
		[]string{"f", "e", "d", "c", "b", "a"}},

	{"test_01",
		&linkedlist.Node[string]{Val: "x",
			Next: &linkedlist.Node[string]{Val: "y"}},
		[]string{"y", "x"}},

	{"test_02",
		&linkedlist.Node[string]{Val: "p"},
		[]string{"p"}},
}

func TestReverseList(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := reverseList(test.input)
			if !reflect.DeepEqual(result.GetValues(), test.expected) {
				t.Errorf("For input '%s', expected result is '%v' but got '%v'", test.input.String(), test.expected, result.GetValues())
			}
		})
	}
}
