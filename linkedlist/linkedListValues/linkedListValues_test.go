package linkedListValues

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
					Next: &linkedlist.Node[string]{Val: "d"}}}},
		[]string{"a", "b", "c", "d"}},

	{"test_01",
		&linkedlist.Node[string]{Val: "x",
			Next: &linkedlist.Node[string]{Val: "y"}},
		[]string{"x", "y"}},

	{"test_02",
		&linkedlist.Node[string]{Val: "q"},
		[]string{"q"}},

	{"test_03",
		nil,
		[]string{}},
}

func TestLinkedListValues(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := linkedListValues(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For input '%s', expected result is '%v' but got '%v'", test.input.String(), test.expected, result)
			}
		})
	}
}
