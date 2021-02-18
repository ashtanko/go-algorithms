package is_subsequence

import "github.com/ashtanko/go-algorithms/utils"

func isSubsequence(source string, target string) bool {
	rows, cols := len(source), len(target)
	if rows == 0 {
		return true
	}
	var matrix = make([][]int, rows+1)
	for i := 0; i < rows+1; i++ {
		matrix[i] = make([]int, cols+1)
	}

	for col := 1; col <= cols; col++ {
		for row := 1; row <= rows; row++ {
			if source[row-1] == target[col-1] {
				matrix[row][col] = matrix[row-1][col-1] + 1
			} else {
				matrix[row][col] = utils.Max(matrix[row][col-1], matrix[row-1][col])
			}
		}
		if matrix[rows][col] == rows {
			return true
		}
	}
	return false
}
