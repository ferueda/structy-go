package longestPath

import (
	"testing"

	"github.com/ferueda/structy-go/graph"
)

var testTable = []struct {
	name     string
	input    graph.Graph[string]
	expected int
}{
	{"test_00",
		graph.Graph[string]{
			"a": {"c", "b"},
			"b": {"c"},
			"c": {},
		},
		2},

	{"test_01",
		graph.Graph[string]{
			"a": {"c", "b"},
			"b": {"c"},
			"c": {},
			"q": {"r"},
			"r": {"s", "u", "t"},
			"s": {"t"},
			"t": {"u"},
			"u": {},
		},
		4},

	{"test_02",
		graph.Graph[string]{
			"h": {"i", "j", "k"},
			"g": {"h"},
			"i": {},
			"j": {},
			"k": {},
			"x": {"y"},
			"y": {},
		},
		2},

	{"test_03",
		graph.Graph[string]{
			"a": {"b"},
			"b": {"c"},
			"c": {},
			"e": {"f"},
			"f": {"g"},
			"g": {"h"},
			"h": {},
		},
		3},
}

func TestLongestPath(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := longestPath(test.input)
			if result != test.expected {
				t.Errorf("For input '%v', expected result is '%v' but got '%v'", test.input, test.expected, result)
			}
		})
	}
}
