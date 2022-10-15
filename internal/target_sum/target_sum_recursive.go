package target_sum

type TargetSumValueRecursive struct {
	*TargetSumValueBase
}

func (m TargetSumValueRecursive) findTargetSumWays(nums []int, target int) int {
	return recursiveHelper(nums, len(nums)-1, target)
}

func recursiveHelper(nums []int, i, sum int) int {
	if i == -1 {
		if sum == 0 {
			return 1
		}
		return 0
	}
	count := 0
	for sign := -1; sign <= 1; sign += 2 {
		count += recursiveHelper(nums, i-1, sum+sign*nums[i])
	}
	return count
}
