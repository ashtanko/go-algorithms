package num_factored_b_trees

import (
	"math"
	"sort"
)

func numFactoredBinaryTrees(arr []int) int {
	// map
	// key: root node value
	// value: number of binary tree
	dp := make(map[int]int)

	constant := (int)(math.Pow(10, 9) + 7)

	// keep A sorted in ascending order
	sort.Ints(arr)

	// scan each possible root node value
	for _, curNum := range arr {

		upperbound := sort.Search(len(arr), func(i int) bool { return arr[i] > int(math.Sqrt(float64(curNum))) })

		// Case 1: cur_num as root with child nodes

		// scan each potential child node value
		for j := 0; j < upperbound; j += 1 {

			factor := arr[j]

			quotient, remainder := curNum/factor, curNum%factor

			// current (factor, quotient) pair are feasible to be child nodes
			if remainder == 0 {

				if quotient == factor {

					// factor is square root of cur_num
					dp[curNum] += dp[factor] * dp[quotient]

				} else {

					// factor is not square root of cur_num
					// (factor, quotient) and (quotient, factor) are symmetric, therefore * 2
					dp[curNum] += dp[factor] * dp[quotient] * 2
				}
			}

		}
		//end of loop j

		// Case 2: cur_num as root without child nodes
		dp[curNum] += 1

	}
	//end of loop i

	totalCount := 0
	for _, count := range dp {
		totalCount += count
	}

	return totalCount % constant
}
