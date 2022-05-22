package findValue

import (
	"testing"

	"github.com/ferueda/structy-go/linkedlist"
)

var intTestTable = []struct {
	name     string
	input1   *linkedlist.Node[int]
	input2   int
	expected bool
}{
	{"test_00",
		&linkedlist.Node[int]{Val: 42},
		42,
		true},

	{"test_01",
		&linkedlist.Node[int]{Val: 42},
		100,
		false},
}

func TestFindIntValue(t *testing.T) {
	for _, test := range intTestTable {
		t.Run(test.name, func(t *testing.T) {
			result := findValue(test.input1, test.input2)
			if result != test.expected {
				t.Errorf("For inputs '%s' and '%v', expected result is '%v' but got '%v'", test.input1.String(), test.input2, test.expected, result)
			}
		})
	}
}

var stringTestTable = []struct {
	name     string
	input1   *linkedlist.Node[string]
	input2   string
	expected bool
}{
	{"test_00",
		&linkedlist.Node[string]{Val: "a",
			Next: &linkedlist.Node[string]{Val: "b",
				Next: &linkedlist.Node[string]{Val: "c",
					Next: &linkedlist.Node[string]{Val: "d"}}}},
		"c",
		true},

	{"test_01",
		&linkedlist.Node[string]{Val: "a",
			Next: &linkedlist.Node[string]{Val: "b",
				Next: &linkedlist.Node[string]{Val: "c",
					Next: &linkedlist.Node[string]{Val: "d"}}}},
		"d",
		true},

	{"test_02",
		&linkedlist.Node[string]{Val: "a",
			Next: &linkedlist.Node[string]{Val: "b",
				Next: &linkedlist.Node[string]{Val: "c",
					Next: &linkedlist.Node[string]{Val: "d"}}}},
		"q",
		false},

	{"test_03",
		&linkedlist.Node[string]{Val: "jason",
			Next: &linkedlist.Node[string]{Val: "leneli"}},
		"jason",
		true},
}

func TestFindStringValue(t *testing.T) {
	for _, test := range stringTestTable {
		t.Run(test.name, func(t *testing.T) {
			result := findValue(test.input1, test.input2)
			if result != test.expected {
				t.Errorf("For inputs '%s' and '%v', expected result is '%v' but got '%v'", test.input1.String(), test.input2, test.expected, result)
			}
		})
	}
}
