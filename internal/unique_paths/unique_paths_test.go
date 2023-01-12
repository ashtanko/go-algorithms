package unique_paths

import (
	"fmt"
	"reflect"
	"testing"
)

type strategy struct {
	taskStrategy func(m int, n int) int
	name         string
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

	tasks = []strategy{
		{uniquePaths, "unique paths"},
		{uniquePathsSimple, "simple"},
	}
)

func TestUniquePaths(t *testing.T) {
	for _, task := range tasks {
		for _, test := range tests {
			testCaseName := fmt.Sprintf("%s input: %d expected: %d", task.name, test.m, test.expected)
			t.Run(testCaseName, func(t *testing.T) {
				if have := task.taskStrategy(test.m, test.n); !reflect.DeepEqual(have, test.expected) {
					t.Errorf(`input: %+v expected: %+v`, test.expected, have)
				}
			})
		}
	}
}
