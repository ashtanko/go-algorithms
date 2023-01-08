package sort_characters_by_frequency

import (
	"sort"
)

func frequencySort(s string) string {
	b := []byte(s)
	freq := make([]int, 128)
	for _, ch := range b {
		freq[ch]++
	}

	sort.Slice(b, func(i, j int) bool {
		freqI, freqJ := freq[b[i]], freq[b[j]]
		if freqI != freqJ {
			return freqI > freqJ
		}

		return b[i] > b[j]
	})

	return string(b)
}
