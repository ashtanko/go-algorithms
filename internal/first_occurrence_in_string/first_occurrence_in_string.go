package first_occurrence_in_string

// 28. Find the Index of the First Occurrence in a String
// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description/
func strStr(haystack string, needle string) int {
	for i := range haystack {
		if haystack[i] == needle[0] {
			comp := len(needle)
			if len(haystack)-i < len(needle) {
				break
			} else if haystack[i:i+comp] == needle {
				return i
			}
		}
	}
	return -1
}
