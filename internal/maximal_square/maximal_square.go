package maximal_square

import (
	"strconv"

	"github.com/ashtanko/go-algorithms/utils"
)

// https://leetcode.com/problems/maximal-square/
func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}
	size := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			dp[i][j], _ = strconv.Atoi(string(matrix[i][j]))
			if i > 0 && j > 0 && matrix[i][j] == '1' {
				dp[i][j] = utils.Min(utils.Min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
			if dp[i][j] > size {
				size = dp[i][j]
			}
		}
	}
	return size * size
}

func maximalSquare2(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	row := len(matrix)
	col := len(matrix[0])
	cur := make([]int, col)
	pre := 0 // previous row and col, last topleft
	size := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			tmp := cur[j]
			if i == 0 || j == 0 || matrix[i][j] == '0' {
				cur[j] = int(matrix[i][j] - '0')
			} else {
				cur[j] = utils.Min(pre, utils.Min(cur[j], cur[j-1])) + 1
			}
			if size < cur[j] {
				size = cur[j]
			}
			pre = tmp
		}
	}
	return size * size
}
