package average_of_levels

import (
	"github.com/ashtanko/go-algorithms/internal/ds/tree"
)

// averageOfLevels calculates the average value of each level in a binary tree
func averageOfLevels(root *tree.TreeNode) []float64 {
	var ans []float64 // Initialize the slice to store the average values

	// If the root is nil, return an empty slice
	if root == nil {
		return ans
	}

	// Perform a breadth-first search on the tree to get the current level and the next level
	level, next := tree.Bfs([]*tree.TreeNode{root})

	// Calculate the average of the current level and append it to the slice
	ans = append(ans, Avg(level))

	// Loop until there are no more levels in the tree
	for len(next) != 0 {
		// Perform a breadth-first search on the next level
		level, next = tree.Bfs(next)

		// Calculate the average of the next level and append it to the slice
		ans = append(ans, Avg(level))
	}

	// Return the slice containing the average values of each level
	return ans
}

// Avg calculates the average of a slice of integers
func Avg(level []int) float64 {
	n := len(level) // Get the number of integers in the slice

	sum := 0                  // Initialize the sum to 0
	for _, v := range level { // Loop through each integer in the slice
		sum += v // Add the integer to the sum
	}

	// Calculate the average by dividing the sum by the number of integers
	// Convert the result to a float64 value
	avg := (float64(sum)) / (float64(n))

	// Return the average value
	return avg
}
