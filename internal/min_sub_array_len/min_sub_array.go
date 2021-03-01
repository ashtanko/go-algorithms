package min_sub_array_len

import (
	"github.com/ashtanko/go-algorithms/utils"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	ans := math.MaxInt64
	left := 0
	sum := 0

	for i := 0; i < n; i++ {
		sum += nums[i]
		for sum >= target {
			ans = utils.Min(ans, i+1-left)
			sum -= nums[left]
			left++
		}

	}

	if ans != math.MaxInt64 {
		return ans
	} else {
		return 0
	}
}
