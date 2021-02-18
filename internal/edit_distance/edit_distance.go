package edit_distance

import "github.com/ashtanko/go-algorithms/utils"

func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	if n*m == 0 {
		return n + m
	}
	matrix := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		matrix[i] = make([]int, m+1)
	}

	for i := 0; i < n+1; i++ {
		matrix[i][0] = i
	}

	for j := 0; j < m+1; j++ {
		matrix[0][j] = j
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			left := matrix[i-1][j] + 1
			down := matrix[i][j-1] + 1
			leftDown := matrix[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				leftDown += 1
			}
			matrix[i][j] = utils.Min(left, utils.Min(down, leftDown))
		}
	}
	return matrix[n][m]
}
