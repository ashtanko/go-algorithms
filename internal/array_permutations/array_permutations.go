package array_permutations

// permute generates all possible permutations of a given slice of integers.
func permute(nums []int) [][]int {
	// Create a map to track visited elements during backtracking.
	visited := make([]bool, len(nums))

	// Recursive helper function to generate permutations.
	var backtrack func([]int)
	var result [][]int

	backtrack = func(currentPermutation []int) {
		// If the current permutation is of the same length as nums, add it to the result.
		if len(currentPermutation) == len(nums) {
			result = append(result, append([]int{}, currentPermutation...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}

			// Mark the current element as visited.
			visited[i] = true

			// Add the current element to the current permutation.
			currentPermutation = append(currentPermutation, nums[i])

			// Recursively generate permutations with the current element included.
			backtrack(currentPermutation)

			// Remove the last element to backtrack.
			currentPermutation = currentPermutation[:len(currentPermutation)-1]

			// Mark the current element as unvisited for the next iteration.
			visited[i] = false
		}
	}

	backtrack([]int{})
	return result
}
