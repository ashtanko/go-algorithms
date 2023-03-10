package min_cost_connect_points

import (
	"fmt"
	"testing"
)

var (
	tests = []struct {
		points   [][]int
		expected int
	}{
		{
			points:   [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}},
			expected: 20,
		},
		{
			points:   [][]int{{3, 12}, {-2, 5}, {-4, 1}},
			expected: 18,
		},
	}

	tasks = []func(points [][]int) int{minCostConnectPointsUnionFind, minCostConnectPointsPrim}
)

func TestMinCostConnectPoints(t *testing.T) {
	for _, fn := range tasks {
		for _, test := range tests {
			t.Run(fmt.Sprint(test.points, test.expected), func(t *testing.T) {
				if have := fn(test.points); have != test.expected {
					t.Errorf(`input: %+vexpected: %+v`, test.expected, have)
				}
			})
		}
	}
}
