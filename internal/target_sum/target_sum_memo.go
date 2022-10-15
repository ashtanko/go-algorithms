package target_sum

import "math"

type TargetSumValueMemoization struct {
	*TargetSumValueBase
}

func (_ TargetSumValueMemoization) findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	total := 0
	for _, v := range nums {
		total += v
	}
	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, 2*total+1)
		for j := range memo[i] {
			memo[i][j] = math.MinInt64
		}
	}
	return memoizationHelper(nums, total, n, target, memo)
}

func memoizationHelper(nums []int, total, i, sum int, memo [][]int) int {
	j := sum + total
	if j < 0 || j >= len(memo[i]) {
		return 0 // for cases where the target sum is way off and definitely unreachable
	}
	if memo[i][j] != math.MinInt64 {
		return memo[i][j]
	}
	if i == 0 {
		if sum == 0 {
			memo[i][j] = 1
		} else {
			memo[i][j] = 0
		}
	} else {
		memo[i][j] = 0
		for sign := -1; sign <= 1; sign += 2 {
			memo[i][j] += memoizationHelper(nums, total, i-1, sum+sign*nums[i-1], memo)
		}
	}
	return memo[i][j]
}
