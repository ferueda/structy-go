package createLinkedlist

import (
	"reflect"
	"testing"

	"github.com/ferueda/structy-go/linkedlist"
)

var testTable = []struct {
	name     string
	input    []string
	expected *linkedlist.Node[string]
}{
	{"test_00",
		[]string{"h", "e", "y"},
		&linkedlist.Node[string]{Val: "h",
			Next: &linkedlist.Node[string]{Val: "e",
				Next: &linkedlist.Node[string]{Val: "y"}}},
	},

	{"test_01",
		[]string{"1", "7", "1", "8"},
		&linkedlist.Node[string]{Val: "1",
			Next: &linkedlist.Node[string]{Val: "7",
				Next: &linkedlist.Node[string]{Val: "1",
					Next: &linkedlist.Node[string]{Val: "8"}}}},
	},

	{"test_02",
		[]string{"a"},
		&linkedlist.Node[string]{Val: "a"},
	},

	{"test_03",
		[]string{},
		nil,
	},
}

func TestCreateLinkedlist(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := createLinkedlist(test.input)
			if !reflect.DeepEqual(result.GetValues(), test.expected.GetValues()) {
				t.Errorf("For input '%v', expected result is '%s' but got '%s'", test.input, test.expected.String(), result.String())
			}
		})
	}

}
