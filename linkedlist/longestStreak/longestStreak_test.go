package longestStreak

import (
	"testing"

	"github.com/ferueda/structy-go/linkedlist"
)

var testTable = []struct {
	name     string
	input    *linkedlist.Node[int]
	expected int
}{
	{"test_00",
		&linkedlist.Node[int]{Val: 5,
			Next: &linkedlist.Node[int]{Val: 5,
				Next: &linkedlist.Node[int]{Val: 7,
					Next: &linkedlist.Node[int]{Val: 7,
						Next: &linkedlist.Node[int]{Val: 7,
							Next: &linkedlist.Node[int]{Val: 6}}}}}},
		3,
	},

	{"test_01",
		&linkedlist.Node[int]{Val: 3,
			Next: &linkedlist.Node[int]{Val: 3,
				Next: &linkedlist.Node[int]{Val: 3,
					Next: &linkedlist.Node[int]{Val: 3,
						Next: &linkedlist.Node[int]{Val: 9,
							Next: &linkedlist.Node[int]{Val: 9}}}}}},
		4,
	},

	{"test_02",
		&linkedlist.Node[int]{Val: 9,
			Next: &linkedlist.Node[int]{Val: 9,
				Next: &linkedlist.Node[int]{Val: 1,
					Next: &linkedlist.Node[int]{Val: 9,
						Next: &linkedlist.Node[int]{Val: 9,
							Next: &linkedlist.Node[int]{Val: 9}}}}}},
		3,
	},

	{"test_03",
		&linkedlist.Node[int]{Val: 5,
			Next: &linkedlist.Node[int]{Val: 5}},
		2,
	},

	{"test_04",
		&linkedlist.Node[int]{Val: 4},
		1,
	},

	{"test_05",
		nil,
		0,
	},
}

func TestLongestStreak(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := longestStreak(test.input)
			if result != test.expected {
				t.Errorf("For inputs '%s', expected result is '%v' but got '%v'", test.input.String(), test.expected, result)
			}
		})
	}

}
