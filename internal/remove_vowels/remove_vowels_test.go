package remove_vowels

import (
	"fmt"
	"testing"
)

var (
	tests = []struct {
		input    string
		expected string
	}{
		{"leetcodeisacommunityforcoders", "ltcdscmmntyfrcdrs"},
		{"aeiou", ""},
		{"", ""},
	}
	task = []func(s string) string{removeVowels}
)

func TestRemoveVowels(t *testing.T) {
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
