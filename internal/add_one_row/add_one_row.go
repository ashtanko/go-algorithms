package add_one_row

import (
	"github.com/ashtanko/go-algorithms/internal/ds/tree"
)

// Function to add a new row of nodes to a binary tree.
// The new nodes are added as children to the existing nodes in the tree.
// The new nodes will have the value 'v' and will be inserted at the depth 'd'.
func addOneRow(root *tree.TreeNode, v, d int) *tree.TreeNode {
	// If the depth 'd' is 1, create a new node with value 'v' and make the current root as its right child.
	// The new node will be the new root of the tree.
	if d == 1 {
		return &tree.TreeNode{
			Val:   v,
			Left:  root,
			Right: nil,
		}
	}
	// If the depth 'd' is not 1, recursively call the function to insert the new node at the appropriate depth.
	return insertToTree(root, v, d-1, 1)
}

// Function to insert a new node with a given value into a binary tree.
// The function takes the root of the tree, the value of the new node, the current depth, and the next depth as arguments.
// The function recursively traverses the tree and inserts the new node at the appropriate depth.
func insertToTree(root *tree.TreeNode, value, depth, n int) *tree.TreeNode {
	// If the root is nil, return nil.
	if root == nil {
		return nil
	}

	// If the current depth is equal to the next depth, insert the new node as the left and right child of the current root.
	if depth == n {
		// Check if root.Left is not nil before creating a new node.
		if root.Left != nil {
			root.Left = &tree.TreeNode{
				Val:   value,
				Left:  root.Left,
				Right: nil,
			}
		}

		// Check if root.Right is not nil before creating a new node.
		if root.Right != nil {
			root.Right = &tree.TreeNode{
				Val:   value,
				Left:  nil,
				Right: root.Right,
			}
		}
	}

	// Recursively insert the new node into the left subtree.
	root.Left = insertToTree(root.Left, value, depth, n+1)
	// Recursively insert the new node into the right subtree.
	root.Right = insertToTree(root.Right, value, depth, n+1)
	// Return the root of the tree.
	return root
}
