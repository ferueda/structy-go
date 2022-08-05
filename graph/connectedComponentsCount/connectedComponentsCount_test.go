package connectedComponentsCount

import (
	"testing"

	"github.com/ferueda/structy-go/graph"
)

var testTable = []struct {
	name     string
	input    graph.Graph[int]
	expected int
}{
	{"test_00",
		graph.Graph[int]{
			0: {8, 1, 5},
			1: {0},
			5: {0, 8},
			8: {0, 5},
			2: {3, 4},
			3: {2, 4},
			4: {3, 2},
		},
		2},

	{"test_01",
		graph.Graph[int]{
			1: {2},
			2: {1, 8},
			6: {7},
			9: {8},
			7: {6, 8},
			8: {9, 7, 2},
		},
		1},

	{"test_02",
		graph.Graph[int]{
			3: {},
			4: {6},
			6: {4, 5, 7, 8},
			8: {6},
			7: {6},
			5: {6},
			1: {2},
			2: {1},
		},
		3},

	{"test_03",
		graph.Graph[int]{},
		0},

	{"test_04",
		graph.Graph[int]{
			0: {4, 7},
			1: {},
			2: {},
			3: {6},
			4: {0},
			6: {3},
			7: {0},
			8: {},
		},
		5},
}

func TestConnectedComponentsCount(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := connectedComponentsCount(test.input)
			if result != test.expected {
				t.Errorf("For input '%v', expected result is '%v' but got '%v'", test.input, test.expected, result)
			}
		})
	}
}
