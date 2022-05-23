package isUnivalue

import (
	"testing"

	"github.com/ferueda/structy-go/linkedlist"
)

var testTable = []struct {
	name     string
	input    *linkedlist.Node[int]
	expected bool
}{
	{"test_00",
		&linkedlist.Node[int]{Val: 7,
			Next: &linkedlist.Node[int]{Val: 7,
				Next: &linkedlist.Node[int]{Val: 7}}},
		true,
	},

	{"test_01",
		&linkedlist.Node[int]{Val: 7,
			Next: &linkedlist.Node[int]{Val: 7,
				Next: &linkedlist.Node[int]{Val: 4}}},
		false,
	},

	{"test_02",
		&linkedlist.Node[int]{Val: 2,
			Next: &linkedlist.Node[int]{Val: 2,
				Next: &linkedlist.Node[int]{Val: 2,
					Next: &linkedlist.Node[int]{Val: 2,
						Next: &linkedlist.Node[int]{Val: 2}}}}},
		true,
	},

	{"test_03",
		&linkedlist.Node[int]{Val: 2,
			Next: &linkedlist.Node[int]{Val: 2,
				Next: &linkedlist.Node[int]{Val: 2,
					Next: &linkedlist.Node[int]{Val: 3,
						Next: &linkedlist.Node[int]{Val: 2}}}}},
		false,
	},

	{"test_04",
		&linkedlist.Node[int]{Val: 1},
		true,
	},
}

func TestIsUnivalue(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := isUnivalue(test.input)
			if result != test.expected {
				t.Errorf("For input '%s', expected result is '%v' but got '%v'", test.input, test.expected, result)
			}
		})
	}
}
