package target_sum

type TargetSumValueDPWithMaps struct {
	*TargetSumValueBase
}

func (_ TargetSumValueDPWithMaps) findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	dp := map[int]int{0: 1}
	for i := 0; i < n; i++ {
		dp2 := make(map[int]int)
		for currSum, currCount := range dp {
			for sign := -1; sign <= 1; sign += 2 {
				dp2[currSum+sign*nums[i]] += currCount
			}
		}
		dp = dp2
	}
	return dp[target]
}
