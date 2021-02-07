package remove_vowels

import "strings"

func removeVowels(s string) string {
	vowelsMap := map[rune]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}
	builder := strings.Builder{}
	builder.Grow(len(s))
	for _, r := range s {
		if _, ok := vowelsMap[r]; !ok {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}
