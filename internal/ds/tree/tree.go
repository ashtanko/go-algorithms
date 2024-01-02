package tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) List() []int {
	arr := make([]int, t.Size())
	if t == nil {
		return arr
	}
	t.Add(arr, 0)
	return arr
}

func (t *TreeNode) Size() int {
	if t == nil {
		return 0
	} else {
		return t.Left.Size() + 1 + t.Right.Size()
	}
}

func (t *TreeNode) Add(arr []int, i int) int {
	if t == nil {
		return i
	}
	arr[i] = t.Val
	i++
	if t.Left != nil {
		i = t.Left.Add(arr, i)
	}
	if t.Right != nil {
		i = t.Right.Add(arr, i)
	}
	return i
}

// Bfs Breadth-First Search
func Bfs(nodes []*TreeNode) ([]int, []*TreeNode) {
	var values []int
	var next []*TreeNode

	for _, node := range nodes {
		values = append(values, node.Val)
		if node.Left != nil {
			next = append(next, node.Left)
		}
		if node.Right != nil {
			next = append(next, node.Right)
		}
	}

	return values, next
}

// Dfs is an algorithm for traversing or searching tree or graph data structures.
func Dfs(node *TreeNode, l int, r int) int {
	ans := 0
	if node != nil {
		if l <= node.Val && node.Val <= r {
			ans += node.Val
		}
		if l < node.Val {
			ans += Dfs(node.Left, l, r)
		}

		if node.Val < r {
			ans += Dfs(node.Right, l, r)
		}
	}
	return ans
}

func (t *TreeNode) Diagram(top string, root string, bottom string) string {
	if t == nil {
		return fmt.Sprintf("%snull\n", root)
	}
	if t.Left == nil && t.Right == nil {
		return fmt.Sprintf("%s%d\n", root, t.Val)
	} else {
		dr := t.Right.Diagram(
			fmt.Sprintf("%s ", top),
			fmt.Sprintf("%s┌──", top),
			fmt.Sprintf("%s│ ", top),
		)

		dl := t.Left.Diagram(
			fmt.Sprintf("%s│ ", bottom),
			fmt.Sprintf("%s└──", bottom),
			fmt.Sprintf("%s ", bottom),
		)

		return fmt.Sprintf("%s%s%d\n%s", dr, root, t.Val, dl)
	}
}
