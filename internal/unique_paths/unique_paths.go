package unique_paths

func uniquePaths(m int, n int) int {
	store := make([][]int, m+1)
	for i := range store {
		store[i] = make([]int, n+1)
	}
	return helper(m, n, store)
}

func helper(m int, n int, store [][]int) int {
	if store[m][n] != 0 {
		return store[m][n]
	}
	if m == 1 || n == 1 {
		return 1
	}
	store[m][n] = helper(m-1, n, store) + helper(m, n-1, store)
	return store[m][n]
}

func uniquePathsSimple(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = 1
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}

	return dp[n-1]
}
