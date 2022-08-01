package levelAverages

import (
	"reflect"
	"testing"

	"github.com/ferueda/structy-go/binarytree"
)

var testTable = []struct {
	name     string
	input    *binarytree.Node[int]
	expected []float64
}{
	{"test_00",
		&binarytree.Node[int]{Val: 3,
			Left: &binarytree.Node[int]{Val: 11,
				Left:  &binarytree.Node[int]{Val: 4},
				Right: &binarytree.Node[int]{Val: -2}},
			Right: &binarytree.Node[int]{Val: 4,
				Right: &binarytree.Node[int]{Val: 1}}},
		[]float64{3, 7.5, 1}},

	{"test_01",
		&binarytree.Node[int]{Val: 5,
			Left: &binarytree.Node[int]{Val: 11,
				Left: &binarytree.Node[int]{Val: 20},
				Right: &binarytree.Node[int]{Val: 15,
					Left:  &binarytree.Node[int]{Val: 1},
					Right: &binarytree.Node[int]{Val: 3}}},
			Right: &binarytree.Node[int]{Val: 54}},
		[]float64{5, 32.5, 17.5, 2}},

	{"test_02",
		&binarytree.Node[int]{Val: -1,
			Left: &binarytree.Node[int]{Val: -6,
				Left: &binarytree.Node[int]{Val: -3},
				Right: &binarytree.Node[int]{Val: 0,
					Left: &binarytree.Node[int]{Val: -1}}},
			Right: &binarytree.Node[int]{Val: -5,
				Right: &binarytree.Node[int]{Val: 45,
					Right: &binarytree.Node[int]{Val: -2}}}},
		[]float64{-1, -5.5, 14, -1.5}},

	{"test_03",
		&binarytree.Node[int]{Val: 13,
			Left: &binarytree.Node[int]{Val: 4,
				Right: &binarytree.Node[int]{Val: 9,
					Left: &binarytree.Node[int]{Val: 2,
						Left: &binarytree.Node[int]{Val: 42}}}},
			Right: &binarytree.Node[int]{Val: 2}},
		[]float64{13, 3, 9, 2, 42}},

	{"test_04",
		nil,
		[]float64{}},
}

func TestLevelAverages(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			result := levelAverages(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For inputs '%v', expected result is '%v' but got '%v'", test.input, test.expected, result)
			}
		})
	}
}
