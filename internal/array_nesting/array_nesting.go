package array_nesting

func arrayNesting(nums []int) int {
	maxSize := 0
	for i := 0; i < len(nums); i++ {
		size := 0
		for k := i; nums[k] != -1; size++ {
			k, nums[k] = nums[k], -1
		}
		if size > maxSize {
			maxSize = size
		}
	}
	return maxSize
}
