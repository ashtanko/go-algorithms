package min_cost_climbing_stairs

import "github.com/ashtanko/go-algorithms/utils"

func minCostClimbingStairs(cost []int) int {
	f1, f2 := 0, 0
	for i := len(cost) - 1; i >= 0; i-- {
		f0 := cost[i] + utils.Min(f1, f2)
		f2 = f1
		f1 = f0
	}
	return utils.Min(f1, f2)
}
