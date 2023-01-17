package group_anagrams

import (
	"sort"
	"strings"
)

// 49. Group Anagrams
// https://leetcode.com/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
	var result [][]string
	var ref = make(map[string]int)

	for k, v := range strs {
		s := strings.Split(v, "")
		sort.Strings(s)
		str := strings.Join(s, "")
		index, ok := ref[str]
		if !ok {
			result = append(result, []string{strs[k]})
			ref[str] = len(result) - 1
			continue
		}

		result[index] = append(result[index], strs[k])
	}

	return result
}
