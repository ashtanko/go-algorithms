package target_sum

type TargetSumValueBacktracking struct {
	*TargetSumValueBase
}

func (_ TargetSumValueBacktracking) findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	prefixSums := make([]int, n)
	prefixSums[0] = nums[0]
	for i := 1; i < n; i++ {
		prefixSums[i] = prefixSums[i-1] + nums[i]
	}
	return backtrackingHelper(nums, n-1, target, prefixSums)
}

func backtrackingHelper(nums []int, i, sum int, prefixSums []int) int {
	if i == -1 {
		if sum == 0 {
			return 1
		}
		return 0
	}

	// prune subtrees that are not promising in O(1) with the help of the precomputed prefixSums.
	if sum+prefixSums[i] < 0 || sum-prefixSums[i] > 0 {
		return 0
	}

	count := 0
	for sign := -1; sign <= 1; sign += 2 {
		count += backtrackingHelper(nums, i-1, sum+sign*nums[i], prefixSums)
	}
	return count
}
