package target_sum

type TargetSumValue2DDP struct {
	*TargetSumValueBase
}

func (_ TargetSumValue2DDP) findTargetSumWays(nums []int, target int) int {
	print("Where is the panic: ", nums, target)
	n := len(nums)
	total := 0
	for _, v := range nums {
		total += v
	}
	m := 2*total + 1

	// terminate early if target sum is definitely unreachable
	if j := target + total; j < 0 || j >= m {
		return 0
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	dp[0][0+total] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if dp[i][j] > 0 {
				for sign := -1; sign <= 1; sign += 2 {
					if nextSum := j + sign*nums[i]; nextSum >= 0 && nextSum < m {
						dp[i+1][nextSum] += dp[i][j]
					}
				}
			}
		}
	}
	return dp[n][target+total]
}
