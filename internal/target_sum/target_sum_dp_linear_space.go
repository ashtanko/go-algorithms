package target_sum

type TargetSumValueDPLinearSpace struct {
	*TargetSumValueBase
}

func (_ TargetSumValueDPLinearSpace) findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	total := 0
	for _, v := range nums {
		total += v
	}
	m := 2*total + 1

	if j := target + total; j < 0 || j >= m {
		return 0
	}

	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	dp[0][0+total] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if dp[0][j] > 0 {
				for sign := -1; sign <= 1; sign += 2 {
					if nextSum := j + sign*nums[i]; nextSum >= 0 && nextSum < m {
						dp[1][nextSum] += dp[0][j]
					}
				}
			}
		}
		dp[0], dp[1] = dp[1], dp[0] // swap the two rows
		for j := 0; j < m; j++ {
			dp[1][j] = 0 // clear the new second row for the next iteration
		}
	}
	return dp[0][target+total]
}
