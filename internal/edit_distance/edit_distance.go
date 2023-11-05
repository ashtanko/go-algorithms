package edit_distance

import "github.com/ashtanko/go-algorithms/utils"

// minDistance calculates the minimum number of operations to transform word1 into word2
func minDistance(word1 string, word2 string) int {
	lenWord1, lenWord2 := len(word1), len(word2)

	// If either word is empty, return the length of the other word
	if lenWord1*lenWord2 == 0 {
		return lenWord1 + lenWord2
	}

	// Initialize the matrix with zeros
	matrix := make([][]int, lenWord1+1)
	for i := 0; i <= lenWord1; i++ {
		matrix[i] = make([]int, lenWord2+1)
		for j := 0; j <= lenWord2; j++ {
			matrix[i][j] = 0
		}
	}

	// Initialize the first row and column of the matrix
	for i := 0; i <= lenWord1; i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= lenWord2; j++ {
		matrix[0][j] = j
	}

	// Fill the rest with the matrix
	for i := 1; i <= lenWord1; i++ {
		for j := 1; j <= lenWord2; j++ {
			insert := matrix[i-1][j] + 1
			deleted := matrix[i][j-1] + 1
			substitute := matrix[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				substitute += 1
			}
			matrix[i][j] = utils.Min(insert, utils.Min(deleted, substitute))
		}
	}

	// Return the minimum number of operations
	return matrix[lenWord1][lenWord2]
}
