package increasing_triplet

import "math"

// 334. Increasing Triplet Subsequence
// https://leetcode.com/problems/increasing-triplet-subsequence/
func increasingTriplet(nums []int) bool {
	if nums == nil || len(nums) < 3 {
		return false
	}
	first := nums[0]
	second := math.MaxInt16

	for i := 1; i < len(nums); i++ {
		if nums[i] <= first {
			first = nums[i]
		} else if second != math.MaxInt16 && nums[i] > second {
			return true
		} else {
			second = nums[i]
		}
	}
	return false
}
