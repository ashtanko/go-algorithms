package valid_parentheses

import (
	"fmt"
	"testing"
)

var (
	tests = []struct {
		s    string
		want bool
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
			t.Run(fmt.Sprint(test.s, test.want), func(t *testing.T) {
				t.Parallel()
				if have := fn(test.s); have != test.want {
					t.Errorf(`
want: %+v
have: %+v`, test.want, have)
				}
			})
		}
	}
}
