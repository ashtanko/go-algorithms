package valid_parentheses

import (
	"fmt"
	"testing"
)

var (
	tests = []struct {
		input    string
		expected bool
	}{
		{"", false},
		{"(((((())))))", true},
		{"()()()()", true},
		{"(((((((()", false},
		{"))((", false},
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"]", false},
	}
	task = []func(s string) bool{isValidParentheses}
)

func TestValidParentheses(t *testing.T) {
	for _, fn := range task {
		for _, test := range tests {
			test := test
			t.Run(fmt.Sprint(test.input, test.expected), func(t *testing.T) {
				t.Parallel()
				if have := fn(test.input); have != test.expected {
					t.Errorf(`input: %+vexpected: %+v`, test.expected, have)
				}
			})
		}
	}
}
