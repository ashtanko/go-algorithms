package maximal_square

import (
	"fmt"
	"testing"
)

type strategy struct {
	task func(matrix [][]byte) int
	name string
}

var (
	tests = []struct {
		matrix   [][]byte
		expected int
	}{
		{
			[][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			4,
		},
		{
			[][]byte{
				{'0', '1'},
				{'1', '0'},
			},
			1,
		},
		{
			[][]byte{
				{'0'},
			},
			0,
		},
	}

	task = []strategy{{maximalSquare, "str1"}, {maximalSquare2, "srt2"}}
)

func TestGetPermutation(t *testing.T) {
	for _, fn := range task {
		for _, test := range tests {
			testCase := fmt.Sprintf("%s %d %d", fn.name, test.matrix, test.expected)
			t.Run(testCase, func(t *testing.T) {
				t.Parallel()
				if have := fn.task(test.matrix); have != test.expected {
					t.Errorf(`input: %+vexpected: %+v`, test.expected, have)
				}
			})
		}
	}
}
