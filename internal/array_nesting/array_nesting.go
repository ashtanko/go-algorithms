package array_nesting

// arrayNesting finds the size of the longest cycle in an array where elements form cycles by following their values as indices.
func arrayNesting(nums []int) int {
	maxSize := 0

	// Iterate through the elements in the input slice
	for i := 0; i < len(nums); i++ {
		size := 0

		// Follow the cycle by repeatedly updating indices and marking visited elements as -1
		for k := i; nums[k] != -1; size++ {
			k, nums[k] = nums[k], -1
		}

		// Update maxSize if the current cycle size is larger
		if size > maxSize {
			maxSize = size
		}
	}

	return maxSize
}
