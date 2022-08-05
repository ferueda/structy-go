package hasPath

import (
	"testing"

	"github.com/ferueda/structy-go/graph"
)

var testTable = []struct {
	name     string
	input1   graph.Graph[string]
	input2   string
	input3   string
	expected bool
}{
	{"test_00",
		graph.Graph[string]{
			"f": {"g", "i"},
			"g": {"h"},
			"h": {},
			"i": {"g", "k"},
			"j": {"i"},
			"k": {},
		},
		"f",
		"k",
		true},

	{"test_01",
		graph.Graph[string]{
			"f": {"g", "i"},
			"g": {"h"},
			"h": {},
			"i": {"g", "k"},
			"j": {"i"},
			"k": {},
		},
		"f",
		"j",
		false},

	{"test_02",
		graph.Graph[string]{
			"f": {"g", "i"},
			"g": {"h"},
			"h": {},
			"i": {"g", "k"},
			"j": {"i"},
			"k": {},
		},
		"i",
		"h",
		true},

	{"test_03",
		graph.Graph[string]{
			"v": {"x", "w"},
			"w": {},
			"x": {},
			"y": {"z"},
			"z": {},
		},
		"v",
		"w",
		true},

	{"test_04",
		graph.Graph[string]{
			"v": {"x", "w"},
			"w": {},
			"x": {},
			"y": {"z"},
			"z": {},
		},
		"v",
		"z",
		false},
}

func TestHasPath(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := hasPath(test.input1, test.input2, test.input3)
			if result != test.expected {
				t.Errorf("For inputs '%v', '%v', and '%v', expected result is '%v' but got '%v'", test.input1, test.input2, test.input3, test.expected, result)
			}
		})
	}
}
