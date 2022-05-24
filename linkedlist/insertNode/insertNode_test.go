package insertNode

import (
	"reflect"
	"testing"

	"github.com/ferueda/structy-go/linkedlist"
)

var testTable = []struct {
	name     string
	input1   *linkedlist.Node[string]
	input2   string
	input3   int
	expected []string
}{
	{"test_00",
		&linkedlist.Node[string]{Val: "a",
			Next: &linkedlist.Node[string]{Val: "b",
				Next: &linkedlist.Node[string]{Val: "c",
					Next: &linkedlist.Node[string]{Val: "d"}}}},
		"x",
		2,
		[]string{"a", "b", "x", "c", "d"}},

	{"test_01",
		&linkedlist.Node[string]{Val: "a",
			Next: &linkedlist.Node[string]{Val: "b",
				Next: &linkedlist.Node[string]{Val: "c",
					Next: &linkedlist.Node[string]{Val: "d"}}}},
		"v",
		3,
		[]string{"a", "b", "c", "v", "d"}},

	{"test_02",
		&linkedlist.Node[string]{Val: "a",
			Next: &linkedlist.Node[string]{Val: "b",
				Next: &linkedlist.Node[string]{Val: "c",
					Next: &linkedlist.Node[string]{Val: "d"}}}},
		"m",
		4,
		[]string{"a", "b", "c", "d", "m"}},

	{"test_03",
		&linkedlist.Node[string]{Val: "a",
			Next: &linkedlist.Node[string]{Val: "b"}},
		"z",
		0,
		[]string{"z", "a", "b"}},
}

func TestInsertNode(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := insertNode(test.input1, test.input2, test.input3)
			if !reflect.DeepEqual(result.GetValues(), test.expected) {
				t.Errorf("For inputs '%s', '%s', and '%v', expected result is '%v' but got '%v'", test.input1.String(), test.input2, test.input3, test.expected, result.GetValues())
			}
		})
	}

}
