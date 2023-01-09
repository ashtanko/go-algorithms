package unique_paths

import (
	"fmt"
	"reflect"
	"testing"
)

type strategy struct {
	task func(m int, n int) int
	name string
}

var (
	tests = []struct {
		m        int
		n        int
		expected int
	}{
		{
			3,
			7,
			28,
		},
		{
			3,
			2,
			3,
		},
	}

	task = []strategy{
		{uniquePaths, "unique paths"},
		{uniquePathsSimple, "unique paths simple"},
	}
)

func TestUniquePaths(t *testing.T) {
	for _, fn := range task {
		for _, test := range tests {
			testCaseName := fmt.Sprintf("%s input: %d expected: %d", fn.name, test.m, test.expected)
			t.Run(testCaseName, func(t *testing.T) {
				t.Parallel()
				if have := fn.task(test.m, test.n); !reflect.DeepEqual(have, test.expected) {
					t.Errorf(`input: %+v expected: %+v`, test.expected, have)
				}
			})
		}
	}
}
