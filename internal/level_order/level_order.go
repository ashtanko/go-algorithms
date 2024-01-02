package level_order

import (
	"github.com/ashtanko/go-algorithms/internal/ds/tree"
)

func levelOrder(root *tree.TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}

	level, next := tree.Bfs([]*tree.TreeNode{root})
	ans = append(ans, level)

	for len(next) != 0 {
		level, next = tree.Bfs(next)
		ans = append(ans, level)
	}
	return ans
}
