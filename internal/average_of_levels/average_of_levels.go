package average_of_levels

import (
	"github.com/ashtanko/go-algorithms/ds/tree"
)

func averageOfLevels(root *tree.TreeNode) []float64 {
	var ans []float64
	if root == nil {
		return ans
	}

	level, next := tree.Bfs([]*tree.TreeNode{root})

	ans = append(ans, avg(level))

	for len(next) != 0 {
		level, next = tree.Bfs(next)
		ans = append(ans, avg(level))
	}
	return ans
}

func avg(level []int) float64 {
	n := len(level)
	sum := 0
	for _, v := range level {
		sum += v
	}
	avg := (float64(sum)) / (float64(n))
	return avg
}
