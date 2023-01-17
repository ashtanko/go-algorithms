package sum_numbers

import "github.com/ashtanko/go-algorithms/ds/tree"

// 129. Sum Root to Leaf Numbers
func sumNumbersDFS(root *tree.TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *tree.TreeNode, cur int) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		return cur*10 + node.Val
	}
	return dfs(node.Left, cur*10+node.Val) + dfs(node.Right, cur*10+node.Val)
}

func sumNumbers(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	queueNode, queueVal, ans := []*tree.TreeNode{root}, []int{root.Val}, 0
	for len(queueNode) > 0 {
		curNode, curVal := queueNode[0], queueVal[0]
		queueNode, queueVal = queueNode[1:], queueVal[1:]
		if curNode.Left == nil && curNode.Right == nil {
			ans += curVal
		}
		if curNode.Left != nil {
			queueNode = append(queueNode, curNode.Left)
			queueVal = append(queueVal, curVal*10+curNode.Left.Val)
		}
		if curNode.Right != nil {
			queueNode = append(queueNode, curNode.Right)
			queueVal = append(queueVal, curVal*10+curNode.Right.Val)
		}
	}
	return ans
}
