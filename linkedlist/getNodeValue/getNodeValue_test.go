package getNodeValue

import (
	"testing"

	"github.com/ferueda/structy-go/linkedlist"
)

var testTable = []struct {
	name     string
	input1   *linkedlist.Node[string]
	input2   int
	expected string
}{
	{"test_00",
		&linkedlist.Node[string]{Val: "a",
			Next: &linkedlist.Node[string]{Val: "b",
				Next: &linkedlist.Node[string]{Val: "c",
					Next: &linkedlist.Node[string]{Val: "d"}}}},
		2,
		"c"},

	{"test_01",
		&linkedlist.Node[string]{Val: "a",
			Next: &linkedlist.Node[string]{Val: "b",
				Next: &linkedlist.Node[string]{Val: "c",
					Next: &linkedlist.Node[string]{Val: "d"}}}},
		3,
		"d"},

	{"test_02",
		&linkedlist.Node[string]{Val: "a",
			Next: &linkedlist.Node[string]{Val: "b",
				Next: &linkedlist.Node[string]{Val: "c",
					Next: &linkedlist.Node[string]{Val: "d"}}}},
		7,
		""},

	{"test_03",
		&linkedlist.Node[string]{Val: "banana",
			Next: &linkedlist.Node[string]{Val: "mango"}},
		0,
		"banana"},

	{"test_04",
		&linkedlist.Node[string]{Val: "banana",
			Next: &linkedlist.Node[string]{Val: "mango"}},
		1,
		"mango"},
}

func TestGetNodeValue(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := getNodeValue(test.input1, test.input2)
			if result != test.expected {
				t.Errorf("For inputs '%s' and '%v', expected result is '%v' but got '%v'", test.input1.String(), test.input2, test.expected, result)
			}
		})
	}
}
