package word_break

// TopDown DP without Memoization
func wordBreakTopDown(s string, wordDict []string) bool {
	dict := map[string]bool{}
	for _, word := range wordDict {
		dict[word] = true
	}

	var dfs func(s string) bool
	dfs = func(s string) bool {
		if len(s) == 0 {
			return true
		}
		for i := 1; i <= len(s); i++ {
			if dict[s[:i]] && dfs(s[i:]) {
				return true
			}
		}
		return false
	}
	return dfs(s)
}

// TopDown DP with Memoization
func wordBreakMemo(s string, wordDict []string) bool {
	dict := map[string]bool{}
	for _, word := range wordDict {
		dict[word] = true
	}

	memo := map[string]bool{}
	var dfs func(s string) bool
	dfs = func(s string) bool {
		if len(s) == 0 {
			return true
		}
		if canBreak, found := memo[s]; found {
			return canBreak
		}
		for i := 1; i <= len(s); i++ {
			if dict[s[:i]] && dfs(s[i:]) {
				memo[s] = true
				return true
			}
		}
		memo[s] = false
		return false
	}
	return dfs(s)
}

// BottomUp DP 1D
func wordBreak1D(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	for _, word := range wordDict {
		dict[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dict[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
