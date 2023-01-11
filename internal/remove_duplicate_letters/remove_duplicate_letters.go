package remove_duplicate_letters

// 316. Remove Duplicate Letters
// https://leetcode.com/problems/remove-duplicate-letters/
func removeDuplicateLetters(s string) string {
	stock := make(map[rune]int)
	for _, r := range s {
		stock[r]++
	}
	visited := make(map[rune]bool)
	out := ""
	for _, r := range s {
		if visited[r] {
			stock[r]--
			continue
		}
		for out != "" && rune(out[len(out)-1]) > r && stock[rune(out[len(out)-1])] != 0 {
			visited[rune(out[len(out)-1])] = false
			out = out[:len(out)-1]
		}
		out += string(r)
		visited[r] = true
		stock[r]--
	}
	return out
}
